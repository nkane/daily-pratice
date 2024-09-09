package leetcode

import (
	"testing"

	"gotest.tools/assert"
)

func TestFindNumbers(t *testing.T) {
	tests := []struct {
		Name   string
		Nums   []int
		Output int
	}{
		{
			Name:   "Example 1",
			Nums:   []int{12, 345, 2, 6, 7896},
			Output: 2,
		},
		{
			Name:   "Example 2",
			Nums:   []int{555, 901, 482, 1771},
			Output: 1,
		},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			got := findNumbers(test.Nums)
			assert.Assert(t, test.Output == got)
		})
	}
}

func TestSortedSquares(t *testing.T) {
	tests := []struct {
		Name   string
		Nums   []int
		Output []int
	}{
		{
			Name:   "Example 1",
			Nums:   []int{-4, -1, 0, 3, 10},
			Output: []int{0, 1, 9, 16, 100},
		},
		{
			Name:   "Example 2",
			Nums:   []int{-7, -3, 2, 3, 11},
			Output: []int{4, 9, 9, 49, 121},
		},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			got := sortedSquares(test.Nums)
			assert.DeepEqual(t, test.Output, got)
		})
	}
}
