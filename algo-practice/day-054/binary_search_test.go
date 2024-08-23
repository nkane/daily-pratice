package practice

import (
	"testing"

	"gotest.tools/assert"
)

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		A             []int
		Target        int
		ExpectedIdx   int
		ExpectedCount int
	}{
		{
			A:             []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
			Target:        8,
			ExpectedIdx:   7,
			ExpectedCount: 1,
		},
	}
	for _, test := range tests {
		idx, count := BinarySearch(test.A, test.Target)
		assert.Assert(t, test.ExpectedIdx == idx)
		assert.Assert(t, test.ExpectedCount == count)
	}
}

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
