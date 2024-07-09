package practice

import (
	"sort"
	"testing"
)

func TestQuicksort(t *testing.T) {
	a := []int{1024, 2, 64, 256, 128, 4, 32, 16, 32, 0}
	Quicksort(a, 0, len(a)-1)
	if !sort.IntsAreSorted(a) {
		t.FailNow()
	}
}
