package leetcode

import (
	"testing"

	"gotest.tools/assert"
)

func TestFindCircleNum(t *testing.T) {
	tests := []struct {
		Name   string
		Input  [][]int
		Output int
	}{
		{
			Name: "Two Providences",
			Input: [][]int{
				{1, 1, 0},
				{1, 1, 0},
				{0, 0, 1},
			},
			Output: 2,
		},
		{
			Name: "Three Providences",
			Input: [][]int{
				{1, 0, 0},
				{0, 1, 0},
				{0, 0, 1},
			},
			Output: 3,
		},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			got := findCircleNum(test.Input)
			assert.Assert(t, test.Output == got, "expected: %+v, got %+v", test.Output, got)
		})
	}
}
