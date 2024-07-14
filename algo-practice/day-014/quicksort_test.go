package practice

import (
	"sort"
	"testing"

	"gotest.tools/v3/assert"
)

func TestQuickSort(t *testing.T) {
	a := []int{1024, 2, 52, 2, 1, 1234, 9, 0, 1, 2, 9912, 11, -1, 247, 22, 99}
	Quicksort(a, 0, len(a)-1)
	assert.Assert(t, sort.IntsAreSorted(a) == true)
}
