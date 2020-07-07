package hg

import (
	"os"
	"os/exec"
)

// Clone makes a child process and runs `git clone repoURL`.
// The cloned repo is stored at the path.
func Clone(repoURL, path string) error {
	// Setup the command.
	cmd := exec.Command("hg", "clone", repoURL, path)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	// And execute it.
	return cmd.Run()
}
