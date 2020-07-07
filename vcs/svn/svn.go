package svn

import (
	"os"
	"os/exec"
)

// Clone makes a child process and runs `svn checkout repoURL`.
// The cloned repo is stored at the path.
func Clone(repoURL, path string) error {
	// Setup the command.
	cmd := exec.Command("svn", "checkout", repoURL, path)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	// And execute it.
	return cmd.Run()
}
