package leetcode

import (
	"testing"

	"gotest.tools/assert"
)

func TestProductExceptSelf(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	expected := []int{24, 12, 8, 6}

	got := productExceptSelf(nums)
	assert.DeepEqual(t, expected, got)

	nums = []int{-1, 1, 0, -3, 3}
	expected = []int{0, 0, 9, 0, 0}
	got = productExceptSelf(nums)
	assert.DeepEqual(t, expected, got)

	nums = []int{1, 2}
	got = productExceptSelf(nums)
}

func TestProductExceptSelfSlow(t *testing.T) {
	nums := []int{4, 5, 1, 8, 2, 10, 6}
	expected := []int{4800, 3840, 19200, 2400, 9600, 1920, 3200}
	got := productExpectSelfSlower(nums)
	assert.DeepEqual(t, expected, got)
}
