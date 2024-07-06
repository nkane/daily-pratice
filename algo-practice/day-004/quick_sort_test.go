package practice

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestQuickSort(t *testing.T) {
	a := []int{10, 2, 5, 6, 1, 8}
	expected := []int{1, 2, 5, 6, 8, 10}

	Quicksort(a, 0, len(a)-1)

	assert.DeepEqual(t, expected, a)
}
