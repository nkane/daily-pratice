package practice

import (
	"sort"
	"testing"

	"gotest.tools/assert"
)

func TestQuicksort(t *testing.T) {
	a := []int{32, 64, 2, 128, 4, 8, 256, 16, 1024, 0}
	expected := []int{32, 64, 2, 128, 4, 8, 256, 16, 1024, 0}
	sort.Ints(expected)
	Quicksort[int](a, 0, len(a)-1)
	assert.DeepEqual(t, expected, a)
}
