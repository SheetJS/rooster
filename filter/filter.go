package filter

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// ExtensionMap maps a file extension to a list of files that
// match that extension.
type ExtensionMap map[string][]string

// Find finds all files in root that match any extensions in
// extensions.
func Find(root string, extensions []string) (ExtensionMap, error) {
	s := makeSet(extensions)
	m := make(map[string][]string)

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Get the current file extension from the file.
		fileExtension := strings.ToLower(filepath.Ext(path))

		// See if we can find the file's extension in the hash set.
		if _, in := s[fileExtension]; !in {

			// If we can't find it
			// we'll have to check each extension
			// in the set (suboptimal!), and check for the patter expression.
			// In this case we'll remove the added dot from the string in the hash set.
			//
			// TODO: Optimize this.
			var (
				matched bool
				err     error
			)

			for setExtension := range s {
				matched, err = filepath.Match(setExtension[1:], fileExtension)
				if err != nil {
					return err
				}
				if matched {
					// If we matched we can immediately exit.
					break
				}
			}

			// We landed here because either
			// A) We matched and broke early.
			// B) We exhausted the entire set.
			// We can find out which one by checking matched.
			if !matched {
				// Case B, skip this file.
				return nil
			}
			// Case A, continue ahead.
		}

		// fileExtension found or matched!
		m[fileExtension] = append(m[fileExtension], path)
		return nil
	})

	return m, err
}

// ExtensionsFromReader will return a slice of formatted file extensions
// from the given rd.
func ExtensionsFromReader(rd io.Reader) ([]string, error) {
	csvRD := csv.NewReader(rd)
	records, err := csvRD.ReadAll()
	if err != nil {
		return nil, err
	}
	if len(records) != 1 {
		return nil, fmt.Errorf("wrong amount of rows in reader: expected %d got %d", 1, len(records))
	}
	return PrettyExtensions(records[0]), nil
}

// ExtensionsFromString will return a slice of formatted file extensions
// from the given string.
func ExtensionsFromString(str string) ([]string, error) {
	return ExtensionsFromReader(strings.NewReader(str))
}

// PrettyExtensions turns an unprocessed list of extensions
// to a pretty one.
func PrettyExtensions(unprocessed []string) []string {
	records := make([]string, len(unprocessed))
	for i, ext := range unprocessed {
		ext = strings.TrimSpace(ext)
		ext = strings.ToLower(ext)
		if ext[0] != '.' {
			ext = "." + ext
		}
		records[i] = ext
	}
	return records
}

func makeSet(extensions []string) map[string]struct{} {
	s := make(map[string]struct{})
	for _, extension := range extensions {
		s[extension] = struct{}{}
	}
	return s
}
