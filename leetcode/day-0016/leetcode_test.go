package leetcode

import (
	"testing"

	"gotest.tools/assert"
)

func TestLargestAltitude(t *testing.T) {
	gain := []int{-5, 1, 5, 0, -7}
	output := 1
	got := largestAltitude(gain)
	assert.Assert(t, output == got)

	gain = []int{-4, -3, -2, -1, 4, 3, 2}
	output = 0
	got = largestAltitude(gain)
	assert.Assert(t, output == got)
}

func TestLargestAltitude_Solution(t *testing.T) {
	gain := []int{-5, 1, 5, 0, -7}
	output := 1
	got := largestAltitude_Solution(gain)
	assert.Assert(t, output == got)

	gain = []int{-4, -3, -2, -1, 4, 3, 2}
	output = 0
	got = largestAltitude_Solution(gain)
	assert.Assert(t, output == got)
}
