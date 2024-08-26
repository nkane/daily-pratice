package leetcode

import (
	"testing"

	"gotest.tools/assert"
)

func TestClimbingStairs(t *testing.T) {
	tests := []struct {
		Name   string
		Cost   []int
		Output int
	}{
		{
			Name:   "Example 1",
			Cost:   []int{10, 15, 20},
			Output: 15,
		},
	}
	for _, test := range tests {
		got := minCostClimbingStairs(test.Cost)
		assert.Assert(t, test.Output == got)
	}
}
