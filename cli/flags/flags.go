package flags

import (
	"flag"
)

// Get grabs the command line flags.
func Get() (configFile string) {
	flag.StringVar(&configFile, "config", ".rooster.yaml", "Rooster's config file")
	flag.Parse()
	return configFile
}
