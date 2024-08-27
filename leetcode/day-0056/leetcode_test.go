package leetcode

import (
	"testing"

	"gotest.tools/assert"
)

func TestRob(t *testing.T) {
	tests := []struct {
		Name   string
		Nums   []int
		Output int
	}{
		{
			Name:   "Example 1",
			Nums:   []int{1, 2, 3, 1},
			Output: 4,
		},
		{
			Name:   "Example 2",
			Nums:   []int{2, 7, 9, 3, 1},
			Output: 12,
		},
	}
	for _, test := range tests {
		got := rob(test.Nums)
		assert.Assert(t, test.Output == got)
	}
}
