package practice

import (
	"sort"
	"testing"

	"gotest.tools/assert"
)

func TestQuicksort(t *testing.T) {
	a := []int{1024, 256, 16, 0, 32, 2, 2048, 128, 64}
	Quicksort(a, 0, len(a)-1)
	assert.Assert(t, sort.IntsAreSorted(a) == true)
}
