package practice

import (
	"testing"

	"gotest.tools/assert"
)

func TestQuicksort(t *testing.T) {
	a := []int{
		5, 8, 1, 4, 2,
	}
	expected := []int{
		1, 2, 4, 5, 8,
	}
	Quicksort(a, 0, len(a)-1)
	assert.DeepEqual(t, expected, a)
}
