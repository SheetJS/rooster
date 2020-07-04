package git

import (
	"os"

	libgit "github.com/go-git/go-git/v5"
)

// Clone effectively runs `git clone repoURL`.
// The cloned repo is stored at the path.
func Clone(repoURL, path string) error {
	// Save the cwd.
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	defer os.Chdir(cwd)

	_, err = libgit.PlainClone(path, false, &libgit.CloneOptions{
		URL:      repoURL,
		Progress: os.Stderr,
	})

	return err
}
