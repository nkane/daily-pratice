package practice

import (
	"testing"

	"gotest.tools/assert"
)

func TestBinarySearch(t *testing.T) {
	list := []int{
		1, 2, 3, 4, 5, 6, 7, 8,
	}

	idx, count := BinarySearch(list, 4)
	assert.Assert(t, idx == 3)
	assert.Assert(t, count == 1)

	idx, count = BinarySearch(list, 2)
	assert.Assert(t, idx == 1)
	assert.Assert(t, count == 2)
}
