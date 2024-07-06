package practice

import (
	"testing"

	"gotest.tools/assert"
)

func TestQuickSort(t *testing.T) {
	a := []int{50, 29, 1, 3, 15, 19, 2, 10, 9}
	expected := []int{1, 2, 3, 9, 10, 15, 19, 29, 50}

	Quicksort(a, 0, len(a)-1)
	assert.DeepEqual(t, expected, a)
}
