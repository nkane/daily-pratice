package leetcode

import (
	"testing"

	"gotest.tools/assert"
)

func TestCanConstruct(t *testing.T) {
	tests := []struct {
		Name       string
		RansomNote string
		Magazine   string
		Output     bool
	}{
		{
			Name:       "Example 1",
			RansomNote: "a",
			Magazine:   "b",
			Output:     false,
		},
		{
			Name:       "Example 2",
			RansomNote: "aa",
			Magazine:   "bb",
			Output:     false,
		},
		{
			Name:       "Example 3",
			RansomNote: "aa",
			Magazine:   "aab",
			Output:     true,
		},
	}

	for _, test := range tests {
		got := canConstruct(test.RansomNote, test.Magazine)
		assert.Assert(t, test.Output == got)
	}
}
