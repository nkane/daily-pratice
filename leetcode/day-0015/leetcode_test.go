package leetcode

import (
	"testing"

	"gotest.tools/assert"
)

func TestLongestSubarray(t *testing.T) {
	nums := []int{1, 1, 0, 1}
	output := 3
	got := longestSubarray(nums)
	assert.Assert(t, output == got)

	nums = []int{0, 1, 1, 1, 0, 1, 1, 0, 1}
	output = 5
	got = longestSubarray(nums)
	assert.Assert(t, output == got)

	nums = []int{1, 1, 1}
	output = 2
	got = longestSubarray(nums)
	assert.Assert(t, output == got)
}

func TestLongestSubarray_Solution(t *testing.T) {
	nums := []int{1, 1, 0, 1}
	output := 3
	got := longestSubarray_Solution(nums)
	assert.Assert(t, output == got)

	nums = []int{0, 1, 1, 1, 0, 1, 1, 0, 1}
	output = 5
	got = longestSubarray_Solution(nums)
	assert.Assert(t, output == got)

	nums = []int{1, 1, 1}
	output = 2
	got = longestSubarray_Solution(nums)
	assert.Assert(t, output == got)
}
