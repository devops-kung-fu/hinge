package lib

import (
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {
	afs := &afero.Afero{Fs: afero.NewMemMapFs()}

	afs.Create("go.mod")
	afs.Create("requirements.txt")

	schedule := Schedule{
		Interval: "daily",
		Time:     "",
		TimeZone: "",
	}
	config, err := Generate(afs, ".", schedule)
	assert.NoError(t, err)
	assert.Equal(t, 2, config.Version)
	assert.Len(t, config.Updates, 2)
	assert.NotNil(t, config)
}

func TestDirectoryParserGitHub(t *testing.T) {
	afs := &afero.Afero{Fs: afero.NewMemMapFs()}

	_ = afs.Mkdir(".github", 0755)
	_ = afs.Mkdir(".github/workflows", 0755)
	_, _ = afs.Create(".github/workflows/ci.yml")
	results := directoryParser(afs, `(.*)\.yml`, ".")

	assert.Len(t, results, 1)
	assert.Equal(t, "/.github/workflows", results[0])
}

func TestDirectoryParserGoMod(t *testing.T) {
	afs := &afero.Afero{Fs: afero.NewMemMapFs()}

	_, _ = afs.Create("go.mod")
	results := directoryParser(afs, `go\.mod`, ".")

	assert.Len(t, results, 1)
	assert.Equal(t, "/", results[0])

	if results[0] != "/" {
		t.Error()
	}
}

func TestRemoveDuplicates(t *testing.T) {
	sliceWithDupes := []string{"1", "1", "2", "2", "3", "3"}
	correctSlice := []string{"1", "2", "3"}
	uniqueSlice := removeDuplicates(sliceWithDupes)

	assert.ElementsMatch(t, correctSlice, uniqueSlice)
}
