package leetcode

import (
	"testing"

	"gotest.tools/assert"
)

func TestCombindationSum(t *testing.T) {
	tests := []struct {
		Name   string
		K      int
		N      int
		Output [][]int
	}{
		{
			Name: "Example 1",
			K:    3,
			N:    7,
			Output: [][]int{
				{1, 2, 4},
			},
		},
		{
			Name: "Example 2",
			K:    3,
			N:    9,
			Output: [][]int{
				{1, 2, 6},
				{1, 3, 5},
				{2, 3, 4},
			},
		},
		{
			Name:   "Example 3",
			K:      4,
			N:      1,
			Output: nil,
		},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			got := combinationSum3(test.K, test.N)
			assert.DeepEqual(t, test.Output, got)
		})
	}
}
