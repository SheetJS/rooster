package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/SheetJS/rooster/cli/flags"
	"github.com/SheetJS/rooster/filter"
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

func main() {

	// Make a temp directory to store the cloned repo.
	tDir, err := ioutil.TempDir(os.TempDir(), "rooster-temp-")
	if err != nil {
		panic(err)
	}
	// And make sure we cleanup after ourselves.
	defer func() {
		if err := os.RemoveAll(tDir); err != nil {
			log.Fatalln(err)
		}
	}()

	// Grab the CLI flags.
	flagSet, err := flags.Get()
	if err != nil {
		log.Fatalln(err)
	}

	// Get the associated vsc cloning function.
	clone, found := funcMap[flagSet.VCS]
	if !found {
		log.Fatalf("unsupported vsc: %s\n", flagSet.VCS)
	}

	// Then clone the repo
	if err := clone(flagSet.RepoURL, tDir); err != nil {
		log.Fatalln(err)
	}

	// Filter out the files.
	m, err := filter.Find(tDir, flagSet.Extensions)
	if err != nil {
		log.Fatalln(err)
	}

	// Write them back to disk.
	if err := m.WriteToPath(flagSet.OutputDir, tDir); err != nil {
		log.Fatalln(err)
	}

}
