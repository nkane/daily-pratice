package leetcode

import (
	"testing"

	"gotest.tools/assert"
)

func TestLongOne(t *testing.T) {
	nums := []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}
	k := 2
	got := longestOne(nums, k)
	assert.Assert(t, got == 6)
}
