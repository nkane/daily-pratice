package practice

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestBinaryTest(t *testing.T) {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	idx, count := BinarySearch[int](list, 8)
	assert.Assert(t, idx == 7)
	assert.Assert(t, count == 1)
	idx, count = BinarySearch[int](list, 4)
	assert.Assert(t, idx == 3)
	assert.Assert(t, count == 2)
}
