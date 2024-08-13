package leetcode

import (
	"testing"

	"gotest.tools/assert"
)

func TestNearestExit(t *testing.T) {
	tests := []struct {
		Name    string
		Input   [][]byte
		Entrace []int
		Output  int
	}{
		{
			Name: "Example 1",
			Input: [][]byte{
				{'+', '+', '.', '+'},
				{'.', '.', '.', '+'},
				{'+', '+', '+', '.'},
			},
			Entrace: []int{1, 2},
			Output:  1,
		},
		{
			Name: "Example 2",
			Input: [][]byte{
				{'+', '+', '+'},
				{'.', '.', '.'},
				{'+', '+', '+'},
			},
			Entrace: []int{1, 0},
			Output:  2,
		},
		{
			Name: "Example 3",
			Input: [][]byte{
				{'.', '+'},
			},
			Entrace: []int{0, 0},
			Output:  -1,
		},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			got := nearestExit(test.Input, test.Entrace)
			assert.Assert(t, test.Output == got)
		})
	}
}
