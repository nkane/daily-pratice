package practice

import (
	"sort"
	"testing"

	"gotest.tools/assert"
)

func TestQuicksort(t *testing.T) {
	a := []int{32, 2, 16, 8, 2, 4, 128, 64, 0}
	expected := []int{32, 2, 16, 8, 2, 4, 128, 64, 0}
	sort.Ints(expected)
	Quicksort(a, 0, len(a)-1)
	assert.DeepEqual(t, expected, a)
}
