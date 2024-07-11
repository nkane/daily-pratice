package practice

import (
	"sort"
	"testing"

	"gotest.tools/v3/assert"
)

func TestQuicksort(t *testing.T) {
	a := []int{1024, 2, 256, 32, 8, 64, 0, 128}
	Quicksort[int](a, 0, len(a)-1)
	assert.Assert(t, sort.IntsAreSorted(a) == true)
}
