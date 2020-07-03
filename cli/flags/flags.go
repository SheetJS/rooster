package flags

import (
	"flag"
	"fmt"
	"strings"

	"github.com/SheetJS/rooster/filter"
)

var supportedVCS = map[string]struct{}{
	"git": {},
	"svn": {},
	"hg":  {},
}

// Set represents the command line
// flags.
type Set struct {
	RepoURL    string
	VCS        string
	OutputDir  string
	Extensions []string
}

// Get grabs the command line flags.
func Get() (*Set, error) {
	pRepo := flag.String("repo", "", "The repo to grab")
	pExt := flag.String("ext", "", "The extension(s) to filter in a comma seperated list")
	pVCS := flag.String("type", "git", "The VCS to use [supported options: git,svn,hg]")
	pOut := flag.String("out", "rooster_output", "The directory to store the output")

	flag.Parse()

	vcs := strings.TrimSpace(strings.ToLower(*pVCS))
	if _, in := supportedVCS[vcs]; !in {
		return nil, fmt.Errorf("unsupported vsc: %s", vcs)
	}

	if *pExt == "" {
		return nil, fmt.Errorf("missing required extension flag")
	}

	extensions, err := filter.ExtensionsFromString(*pExt)
	if err != nil {
		return nil, err
	}

	return &Set{
		RepoURL:    *pRepo,
		VCS:        vcs,
		Extensions: extensions,
		OutputDir:  *pOut,
	}, nil
}
