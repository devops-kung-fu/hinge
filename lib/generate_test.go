package lib

import (
	"reflect"
	"testing"

	"github.com/spf13/afero"
)

func TestDirectoryParserGitHub(t *testing.T) {
	fs := afero.NewMemMapFs()
	_ = fs.Mkdir(".github", 0755)
	_ = fs.Mkdir(".github/workflows", 0755)
	_, _ = fs.Create(".github/workflows/ci.yml")
	results := directoryParser(fs, `(.*)\.yml`, ".")
	if results[0] != "/.github/workflows" {
		t.Error()
	}
}

func TestDirectoryParserGoMod(t *testing.T) {
	fs := afero.NewMemMapFs()
	_, _ = fs.Create("go.mod")
	results := directoryParser(fs, `go\.mod`, ".")
	if results[0] != "/" {
		t.Error()
	}
}

func TestRemoveDuplicates(t *testing.T) {
	sliceWithDupes := []string{"1", "1", "2", "2", "3", "3"}
	correctSlice := []string{"1", "2", "3"}
	uniqueSlice := removeDuplicates(sliceWithDupes)
	if !reflect.DeepEqual(correctSlice, uniqueSlice) {
		t.Error()
	}
}
