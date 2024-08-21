package leetcode

import (
	"testing"

	"gotest.tools/assert"
)

func TestFindPeakElement(t *testing.T) {
	tests := []struct {
		Name   string
		Nums   []int
		Output int
	}{
		{
			Name:   "Example 1",
			Nums:   []int{1, 2, 3, 1},
			Output: 2,
		},
		{
			Name:   "Example 2",
			Nums:   []int{1, 2, 1, 3, 5, 6, 4},
			Output: 5,
		},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			got := findPeakElement(test.Nums)
			assert.DeepEqual(t, test.Output, got)
		})
	}
}
