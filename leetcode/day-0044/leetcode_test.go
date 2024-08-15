package leetcode

import (
	"testing"

	"gotest.tools/assert"
)

func TestFindKthLargest(t *testing.T) {
	tests := []struct {
		Name   string
		Input  []int
		K      int
		Output int
	}{
		{
			Name:   "Example 1",
			Input:  []int{3, 2, 1, 5, 6, 4},
			K:      2,
			Output: 5,
		},
		{
			Name:   "Example 2",
			Input:  []int{3, 2, 3, 1, 2, 4, 5, 5, 6},
			K:      4,
			Output: 4,
		},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			got := findKthLargest(test.Input, test.K)
			assert.Assert(t, got == test.Output, "expected: %d, got: %d", test.Output, got)
		})
	}
}

func TestFindKthLargest_Slice(t *testing.T) {
	tests := []struct {
		Name   string
		Input  []int
		K      int
		Output int
	}{
		{
			Name:   "Example 1",
			Input:  []int{3, 2, 1, 5, 6, 4},
			K:      2,
			Output: 5,
		},
		{
			Name:   "Example 2",
			Input:  []int{3, 2, 3, 1, 2, 4, 5, 5, 6},
			K:      4,
			Output: 4,
		},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			got := findKthLargest_Slice(test.Input, test.K)
			assert.Assert(t, got == test.Output, "expected: %d, got: %d", test.Output, got)
		})
	}
}

func TestFindKthLargest_QuickSelect(t *testing.T) {
	tests := []struct {
		Name   string
		Input  []int
		K      int
		Output int
	}{
		{
			Name:   "Example 1",
			Input:  []int{3, 2, 1, 5, 6, 4},
			K:      2,
			Output: 5,
		},
		{
			Name:   "Example 2",
			Input:  []int{3, 2, 3, 1, 2, 4, 5, 5, 6},
			K:      4,
			Output: 4,
		},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			got := findKthLargest_QuickSelect(test.Input, test.K)
			assert.Assert(t, got == test.Output, "expected: %d, got: %d", test.Output, got)
		})
	}
}
