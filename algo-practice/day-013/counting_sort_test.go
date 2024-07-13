package practice

import (
	"sort"
	"testing"

	"gotest.tools/assert"
)

func TestCountingSort(t *testing.T) {
	a := []int{1024, 128, 2, 0, 8, 256, 64, 32, 16, 0, 2048}
	got := CountingSort(a)
	assert.Assert(t, sort.IntsAreSorted(got) == true)
}
