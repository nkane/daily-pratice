package leetcode

import (
	"testing"

	"gotest.tools/assert"
)

func TestSuggestedProducts(t *testing.T) {
	tests := []struct {
		Name      string
		Intervals [][]int
		Output    int
	}{
		{
			Name: "Example 1",
			Intervals: [][]int{
				{1, 2},
				{2, 3},
				{3, 4},
				{1, 3},
			},
			Output: 1,
		},
		{
			Name: "Example 2",
			Intervals: [][]int{
				{1, 2},
				{1, 2},
				{1, 2},
			},
			Output: 2,
		},
		{
			Name: "Example 3",
			Intervals: [][]int{
				{1, 2},
				{2, 3},
			},
			Output: 0,
		},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			got := eraseOverlapIntervals(test.Intervals)
			assert.Assert(t, test.Output == got)
		})
	}
}
