package lib

import (
	"testing"

	"github.com/spf13/afero"
)

func TestFindFiles(t *testing.T) {
	fs := afero.NewMemMapFs()
	files, err := FindFiles(fs, "/", "(.*)\\.txt")
	if len(files) != 0 {
		t.Error()
	}
	if err != nil {
		t.Error()
	}
}