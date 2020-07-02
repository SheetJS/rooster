package vcs

// CloneFunc should clone the repo found at repoURL and store it at
// path.
type CloneFunc func(repoURL, path string) error
