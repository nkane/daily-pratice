package leetcode

import (
	"testing"

	"gotest.tools/assert"
)

func TestMinEatingSpeed(t *testing.T) {
	tests := []struct {
		Name   string
		Piles  []int
		H      int
		Output int
	}{
		{
			Name:   "Example 1",
			Piles:  []int{3, 6, 7, 11},
			H:      8,
			Output: 4,
		},
		{
			Name:   "Example 2",
			Piles:  []int{30, 11, 23, 4, 20},
			H:      5,
			Output: 30,
		},
		{
			Name:   "Example 3",
			Piles:  []int{30, 11, 23, 4, 20},
			H:      6,
			Output: 23,
		},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			got := minEatintgSpeed(test.Piles, test.H)
			assert.Assert(t, test.Output == got, "expected: %+v, got %+v\n", test.Output, got)
		})
	}
}
