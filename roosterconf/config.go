package roosterconf

import (
	"fmt"
	"io"

	"github.com/SheetJS/rooster/filter"
	"gopkg.in/yaml.v2"
)

var supportedVCS = map[string]struct{}{
	"git": {},
	"svn": {},
	"hg":  {},
}

// Set represents the command line
// flags.
type Set struct {
	RepoURL    string   `yaml:"repo"`
	VCS        string   `yaml:"type"`
	OutputDir  string   `yaml:"out"`
	Extensions []string `yaml:"extensions"`
}

func New(rd io.Reader) (tSets []Set, err error) {
	// Get and parse the yaml.
	err = yaml.NewDecoder(rd).Decode(&tSets)
	if err != nil {
		return nil, err
	}

	// Process the data structure.
	var roosterIndex int
	for i := range tSets {

		// If there's no URL or extensions then this is
		// an invalid YAML file.
		if tSets[i].RepoURL == "" {
			return nil, fmt.Errorf("missing repo url at config entry %d", i+1)
		}
		if len(tSets[i].Extensions) <= 0 {
			return nil, fmt.Errorf("missing extension set at config entry %d", i+1)
		}

		// If there's no output directory
		// specified then set it to rooster_output_xxxx
		if tSets[i].OutputDir == "" {
			tSets[i].OutputDir = fmt.Sprintf("rooster_output_%04x", roosterIndex+1)
			roosterIndex++
		}

		// If there's no vcs specified assume git.
		if tSets[i].VCS == "" {
			tSets[i].VCS = "git"
		}
		// Then check if it's supported.
		if _, in := supportedVCS[tSets[i].VCS]; !in {
			return nil, fmt.Errorf("unsupported vcs at config entry %d: %s", i+1, tSets[i].VCS)
		}

		// Make the extensions pretty so that the filter algorithm can use them.
		tSets[i].Extensions = filter.PrettyExtensions(tSets[i].Extensions)
	}

	return tSets, err
}
