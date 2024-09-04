package leetcode

import (
	"testing"

	"gotest.tools/assert"
)

func TestFizzBuzz(t *testing.T) {
	tests := []struct {
		Name   string
		N      int
		Output []string
	}{
		{
			Name:   "Example 1",
			N:      3,
			Output: []string{"1", "2", "Fizz"},
		},
		{
			Name:   "Example 2",
			N:      5,
			Output: []string{"1", "2", "Fizz", "4", "Buzz"},
		},
		{
			Name:   "Example 3",
			N:      15,
			Output: []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"},
		},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			got := fizzBuzz(test.N)
			assert.DeepEqual(t, test.Output, got)
		})
	}
}

func TestNumberofSteps(t *testing.T) {
	tests := []struct {
		Name   string
		Num    int
		Output int
	}{
		{
			Name:   "Example 1",
			Num:    14,
			Output: 6,
		},
		{
			Name:   "Example 2",
			Num:    8,
			Output: 4,
		},
		{
			Name:   "Example 3",
			Num:    123,
			Output: 12,
		},
	}
	for _, test := range tests {
		got := numberOfSteps(test.Num)
		assert.Assert(t, test.Output == got)
	}
}
