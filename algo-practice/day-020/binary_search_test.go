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

	a = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	idx, count = BinarySearch(a, 12)
	assert.Assert(t, idx == 11)
	assert.Assert(t, count == 2)

	a = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	idx, count = BinarySearch(a, 2)
	assert.Assert(t, idx == 1)
	assert.Assert(t, count == 3)
}
