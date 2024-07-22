package practice

import (
	"testing"

	"gotest.tools/assert"
)

func TestBinarySearch(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	idx, count := BinarySearch(a, 8)
	assert.Assert(t, idx == 7)
	assert.Assert(t, count == 1)
}
