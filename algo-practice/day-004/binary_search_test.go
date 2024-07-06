package practice

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestBinarySearch(t *testing.T) {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8}

	idx, count := BinarySearch(list, 4)

	assert.Assert(t, idx == 3)
	assert.Assert(t, count == 1)

	idx, count = BinarySearch(list, 6)
	assert.Assert(t, idx == 5)
	assert.Assert(t, count == 2)
}
