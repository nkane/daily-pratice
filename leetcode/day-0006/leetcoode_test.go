package leetcode

import (
	"testing"

	"gotest.tools/assert"
)

func TestIncreasignTriplet(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	got := increasingTriplet(nums)
	assert.Assert(t, got == true)

	nums = []int{5, 4, 3, 2, 1}
	got = increasingTriplet(nums)
	assert.Assert(t, got == false)

	nums = []int{2, 1, 5, 0, 4, 6}
	got = increasingTriplet(nums)
	assert.Assert(t, got == true)

	nums = []int{20, 100, 10, 12, 5, 13}
	got = increasingTriplet(nums)
	assert.Assert(t, got == true)
}
