package practice

import (
	"testing"

	"gotest.tools/assert"
)

func TestQuickSort(t *testing.T) {
	list := []int{10, 6, 11, 2, 3, 1}
	expected := []int{1, 2, 3, 6, 10, 11}
	Quicksort(list, 0, len(list)-1)
	assert.DeepEqual(t, expected, list)
}
