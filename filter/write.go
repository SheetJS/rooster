package filter

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// WriteToPath writes the given extension map to the filesystem at path/name.
func (m ExtensionMap) WriteToPath(path string, trimPrefix ...string) error {
	// When we make directories use this mode.
	const dirMode os.FileMode = 0777

	// Make the directories
	rootPath := path
	if err := os.MkdirAll(rootPath, dirMode); err != nil {
		return err
	}

	// Loop through the map.
	for ext, tFiles := range m {

		// Only make directories
		// if there's something in the list of files.
		if len(tFiles) <= 0 {
			continue
		}

		// Then loop through the files.
		for _, fileName := range tFiles {

			// Make the file path with the trimmed prefix
			// if applicable.
			dotlessExt := ext[1:]
			relFileName := fileName
			if len(trimPrefix) >= 1 {
				relFileName = strings.TrimPrefix(fileName, trimPrefix[0])
			}
			newFilePath := fmt.Sprintf("%s/%s/%s", rootPath, dotlessExt, relFileName)

			// Make the files directories
			if err := os.MkdirAll(filepath.Dir(newFilePath), dirMode); err != nil {
				return err
			}
			// Then copy the file.
			if err := copyFile(fileName, newFilePath); err != nil {
				return err
			}

		}
	}

	return nil
}

func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	return out.Close()
}
