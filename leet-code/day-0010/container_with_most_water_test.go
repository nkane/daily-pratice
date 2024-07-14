package leetcode

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestMaxArea(t *testing.T) {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	got := maxArea(height)
	assert.Assert(t, got == 49)

	height = []int{1, 1}
	got = maxArea(height)
	assert.Assert(t, got == 1)
}
