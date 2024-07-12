package practice

import (
	"sort"
	"testing"

	"gotest.tools/assert"
)

func TestQuicksort(t *testing.T) {
	a := []int{1024, 2, 64, 2, 0, 8, 16, 32, 64, 128, 256, 2048, 4}
	Quicksort(a, 0, len(a)-1)
	assert.Assert(t, sort.IntsAreSorted(a))
}
