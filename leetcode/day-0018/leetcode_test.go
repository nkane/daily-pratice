package leetcode

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestPivotIndex(t *testing.T) {
	nums := []int{1, 7, 3, 6, 5, 6}
	idx := pivotIndex(nums)
	assert.Assert(t, idx == 3)

	nums = []int{1, 2, 3}
	idx = pivotIndex(nums)
	assert.Assert(t, idx == -1)

	nums = []int{2, 1, -1}
	idx = pivotIndex(nums)
	assert.Assert(t, idx == 0)
}
