package leetcode

import (
	"testing"

	"gotest.tools/assert"
)

func TestMaxOperations(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	k := 5
	got := maxOperations(nums, k)
	expected := 2
	assert.Assert(t, got == expected, "expected: %d, got: %d", expected, got)

	nums = []int{3, 1, 3, 4, 3}
	k = 6
	got = maxOperations(nums, k)
	expected = 1
	assert.Assert(t, got == expected, "expected: %d, got: %d", expected, got)

	nums = []int{3, 1, 5, 1, 1, 1, 1, 1, 2, 2, 3, 2, 2}
	k = 1
	got = maxOperations(nums, k)
	expected = 0
	assert.Assert(t, got == expected, "expected: %d, got: %d", expected, got)
}
