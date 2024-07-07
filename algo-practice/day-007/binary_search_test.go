package practice

import (
	"testing"

	"gotest.tools/assert"
)

func TestBinarySearch(t *testing.T) {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}

	idx, count := BinarySearch(list, 8)
	assert.Assert(t, idx == 7)
	assert.Assert(t, count == 1)

	idx, count = BinarySearch(list, 4)
	assert.Assert(t, idx == 3)
	assert.Assert(t, count == 2)

	idx, count = BinarySearch(list, 3)
	assert.Assert(t, idx == 2)
	assert.Assert(t, count == 4)

	idx, count = BinarySearch(list, 2)
	assert.Assert(t, idx == 1)
	assert.Assert(t, count == 3)

	idx, count = BinarySearch(list, 1)
	assert.Assert(t, idx == 0)
	assert.Assert(t, count == 4)
}
