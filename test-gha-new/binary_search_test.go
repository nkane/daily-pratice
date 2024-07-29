package practice

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestBinarySearch(t *testing.T) {
	list := []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16,
		17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30,
		31, 32,
	}
	index, count := BinarySearch(list, 32)
	assert.Assert(t, index == 31, "expected 31, got %d", index)
	assert.Assert(t, count == 6, "expected 6, got %d", count)

	index, count = BinarySearch(list, 2)
	assert.Assert(t, index == 1, "expected 1, got %d", index)
	assert.Assert(t, count == 4, "expected 4, got %d", count)

	index, count = BinarySearch(list, 16)
	assert.Assert(t, index == 15, "expected 15, got %d", index)
	assert.Assert(t, count == 1, "expected 1, got %d", count)

	index, count = BinarySearch(list, 8)
	assert.Assert(t, index == 7, "expected 7, got %d", index)
	assert.Assert(t, count == 2, "expected 2, got %d", count)
}
