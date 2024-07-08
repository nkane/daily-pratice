package practice

import (
	"sort"
	"testing"

	"gotest.tools/v3/assert"
)

func TestQuicksort(t *testing.T) {
	a := []int{128, 256, 2, 16, 0, 4, 8, 32}
	Quicksort(a, 0, len(a)-1)
	assert.Assert(t, sort.IntsAreSorted(a))
}
