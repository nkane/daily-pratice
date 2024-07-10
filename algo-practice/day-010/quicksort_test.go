package practice

import (
	"sort"
	"testing"

	"gotest.tools/v3/assert"
)

func TestQuicksort(t *testing.T) {
	a := []int{1024, 2, 256, 8, 32, 128, 64, 2048, 0}
	Quicksort(a, 0, len(a)-1)
	assert.Assert(t, sort.IntsAreSorted(a) == true)
}
