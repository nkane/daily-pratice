package leetcode

import (
	"testing"

	"gotest.tools/assert"
)

func TestFindMaxConsecutiveOnes(t *testing.T) {
	tests := []struct {
		Name   string
		Nums   []int
		Output int
	}{
		{
			Name:   "Example 1",
			Nums:   []int{1, 1, 0, 1, 1, 1},
			Output: 3,
		},
		{
			Name:   "Example 2",
			Nums:   []int{1, 0, 1, 1, 0, 1},
			Output: 2,
		},
	}
	for _, test := range tests {
		t.Run(t.Name(), func(t *testing.T) {
			got := findMaxConsecutiveOnes(test.Nums)
			assert.Assert(t, test.Output == got)
		})
	}
}
