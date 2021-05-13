package lib

import (
	"reflect"

	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	sliceWithDupes := []string{"1", "1", "2", "2", "3", "3"}
	correctSlice := []string{"1", "2", "3"}
	uniqueSlice := removeDuplicates(sliceWithDupes)
	if !reflect.DeepEqual(correctSlice, uniqueSlice) {
		t.Error()
	}
}
