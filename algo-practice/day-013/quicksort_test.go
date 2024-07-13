package practice

import (
	"sort"
	"testing"

	"gotest.tools/assert"
)

func TestQuickSort(t *testing.T) {
	a := []int{1024, 128, 2, 0, 8, 256, 64, 32, 16, 0, 2048}
	Quicksort(a, 0, len(a)-1)
	assert.Assert(t, sort.IntsAreSorted(a) == true)
}
