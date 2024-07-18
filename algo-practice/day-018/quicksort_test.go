package practice

import (
	"testing"

	"gotest.tools/assert"
)

func TestQuicksort(t *testing.T) {
	a := []int{0, 10, 128, 1, 2, 9, 12, 8, 100}
	expected := []int{0, 1, 2, 8, 9, 10, 12, 100, 128}
	Quicksort(a, 0, len(a)-1)
	assert.DeepEqual(t, expected, a)
}
