package leetcode

import (
	"testing"

	"gotest.tools/assert"
)

func TestFindMaxAverageSlidingWindow(t *testing.T) {
	nums := []int{1, 12, -5, -6, 50, 3}
	k := 4
	expected := 12.75
	got := findMaxAverage(nums, k)
	assert.Assert(t, got == expected)
}

func TestFindMaxAverageCumulative(t *testing.T) {
	nums := []int{1, 12, -5, -6, 50, 3}
	k := 4
	expected := 12.75
	got := findMaxAverageCumulative(nums, k)
	assert.Assert(t, got == expected)
}
