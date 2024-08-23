package leetcode

import (
	"testing"

	"gotest.tools/assert"
)

func TestLetterCombindations(t *testing.T) {
	tests := []struct {
		Name   string
		Input  string
		Output []string
	}{
		{
			Name:  "Example 1",
			Input: "23",
			Output: []string{
				"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf",
			},
		},
		{
			Name:   "Example 2",
			Input:  "",
			Output: []string{},
		},
		{
			Name:   "Example 3",
			Input:  "2",
			Output: []string{"a", "b", "c"},
		},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			got := letterCombinations(test.Input)
			assert.DeepEqual(t, test.Output, got)
		})
	}
}
