package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sync"

	"github.com/SheetJS/rooster/cli/flags"
	"github.com/SheetJS/rooster/filter"
	"github.com/SheetJS/rooster/roosterconf"
	"github.com/SheetJS/rooster/vcs"
	"github.com/SheetJS/rooster/vcs/git"
	"github.com/SheetJS/rooster/vcs/hg"
	"github.com/SheetJS/rooster/vcs/svn"
)

var funcMap = map[string]vcs.CloneFunc{
	"git": git.Clone,
	"svn": svn.Clone,
	"hg":  hg.Clone,
}

func cloneRepoAndWrite(r roosterconf.Repository) error {

	// Make a temp directory to store the cloned repo.
	tDir, err := ioutil.TempDir(os.TempDir(), "rooster-temp-")
	if err != nil {
		panic(err)
	}
	// And make sure we cleanup after ourselves.
	defer func() {
		if err := os.RemoveAll(tDir); err != nil {
			log.Printf("failed to cleanup after repo: %s @ %s", r.RepoURL, tDir)
		}
	}()

	// Get the associated vsc cloning function.
	clone, found := funcMap[r.VCS]
	if !found {
		return fmt.Errorf("unsupported vsc: %s", r.VCS)
	}

	// Then clone the repo
	if err := clone(r.RepoURL, tDir); err != nil {
		return fmt.Errorf("failed to clone repo %s: %v", r.RepoURL, err)
	}

	// Filter out the files.
	m, err := filter.Find(tDir, r.Extensions)
	if err != nil {
		return fmt.Errorf("failed to filter files for repo %s: %s", r.RepoURL, err)
	}

	// Write them back to disk.
	if err := m.WriteToPath(r.OutputDir, tDir); err != nil {
		return fmt.Errorf("failed to write files from %s: %v", r.RepoURL, err)
	}

	return nil
}

func main() {

	// Grab the flags.
	configFile := flags.Get()
	configFD, err := os.Open(configFile)
	if err != nil {
		log.Fatalln(err)
	}
	defer configFD.Close()

	// Parse the YAML.
	tRepos, err := roosterconf.New(configFD)
	if err != nil {
		log.Fatalln(err)
	}

	// Grab each repo concurrently.
	wg := new(sync.WaitGroup)
	wg.Add(len(tRepos))

	for _, repo := range tRepos {
		go func(r roosterconf.Repository, pWG *sync.WaitGroup) {
			defer wg.Done()
			if err := cloneRepoAndWrite(r); err != nil {
				log.Println(err)
			}
		}(repo, wg)
	}

	// Wait for all the repositories to finish downloading
	wg.Wait()
}
