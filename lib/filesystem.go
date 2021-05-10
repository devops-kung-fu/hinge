package lib

import (
	"os"
	"regexp"

	"github.com/spf13/afero"
)

// FindFiles - Recursively search for files matching a pattern.
func FindFiles(fs afero.Fs, root string, re string) ([]string, error) {
	libRegEx, e := regexp.Compile(re)
	if e != nil {
		return nil, e
	}
	var files []string
	e = afero.Walk(fs, root, func(filePath string, info os.FileInfo, err error) error {
		if err == nil && libRegEx.MatchString(info.Name()) {
			files = append(files, filePath)
		}
		return nil
	})
	if e != nil {
		return nil, e
	}
	return files, nil
}
