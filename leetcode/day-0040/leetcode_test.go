package leetcode

import (
	"testing"

	"gotest.tools/assert"
)

func TestMinReorder(t *testing.T) {
	tests := []struct {
		Name        string
		Input       int
		Connections [][]int
		Output      int
	}{
		{
			Name:  "Three Reorders",
			Input: 6,
			Connections: [][]int{
				{0, 1},
				{1, 3},
				{2, 3},
				{4, 0},
				{4, 5},
			},
			Output: 3,
		},
		{
			Name:  "Two Reorders",
			Input: 5,
			Connections: [][]int{
				{1, 0},
				{1, 2},
				{3, 2},
				{3, 4},
			},
			Output: 2,
		},
		{
			Name:  "No Reorders",
			Input: 3,
			Connections: [][]int{
				{1, 0},
				{2, 0},
			},
			Output: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			got := minReorder(test.Input, test.Connections)
			assert.Assert(t, test.Output == got, "expected %d, got %d", test.Output, got)
		})
	}
}
