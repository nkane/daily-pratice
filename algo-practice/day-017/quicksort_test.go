package practice

import (
	"sort"
	"testing"

	"gotest.tools/v3/assert"
)

func TestQuicksort(t *testing.T) {
	a := []int{1024, 64, 0, 2, 32, 16, 2048, 128, 256}
	expected := []int{1024, 64, 0, 2, 32, 16, 2048, 128, 256}
	sort.Ints(expected)
	Quicksort(a, 0, len(a)-1)
	assert.DeepEqual(t, expected, a)
}
