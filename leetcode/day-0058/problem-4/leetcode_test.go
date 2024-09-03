package leetcode

import (
	"testing"

	"gotest.tools/assert"
)

func TestFindMinArrowShots(t *testing.T) {
	tests := []struct {
		Name   string
		Points [][]int
		Output int
	}{
		{
			Name: "Example 1",
			Points: [][]int{
				{10, 16},
				{2, 8},
				{1, 6},
				{7, 12},
			},
			Output: 2,
		},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			got := findMinArrowShots(test.Points)
			assert.Assert(t, test.Output == got)
		})
	}
}
