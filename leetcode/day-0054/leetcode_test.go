package leetcode

import (
	"testing"

	"gotest.tools/assert"
)

func TestTribonacci(t *testing.T) {
	tests := []struct {
		Name   string
		N      int
		Output int
	}{
		{
			Name:   "Example 1",
			N:      4,
			Output: 4,
		},
		{
			Name:   "Example 2",
			N:      25,
			Output: 1389537,
		},
	}
	for _, test := range tests {
		got := tribonacci(test.N)
		assert.Assert(t, test.Output == got)
	}
}
