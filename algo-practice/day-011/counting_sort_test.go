package practice

import (
	"sort"
	"testing"

	"gotest.tools/assert"
)

func TestCountSort(t *testing.T) {
	a := []int{1024, 2, 256, 32, 8, 64, 0, 128}
	got := CountSort(a)
	assert.Assert(t, sort.IntsAreSorted(got) == true)
}
