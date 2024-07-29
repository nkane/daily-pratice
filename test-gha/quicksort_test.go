package practice

import (
	"math/rand"
	"sort"
	"testing"

	"gotest.tools/v3/assert"
)

func TestQuickSort(t *testing.T) {
	n := 10
	list := make([]int, n)
	for i := 0; i < n; i++ {
		list[i] = rand.Intn(200)
	}

	expected := make([]int, n)
	copy(expected, list)
	sort.Ints(expected)

	Quicksort(list, 0, len(list)-1)

	t.Logf("expected list: %+v\n", expected)
	t.Logf("got list: %+v\n", list)

	assert.DeepEqual(t, expected, list)
}
