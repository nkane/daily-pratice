package leetcode

import (
	"testing"

	"gotest.tools/assert"
)

func TestSuccessfulPairs(t *testing.T) {
	tests := []struct {
		Name    string
		Spells  []int
		Potions []int
		Success int64
		Output  []int
	}{
		{
			Name:    "Example 1",
			Spells:  []int{5, 1, 3},
			Potions: []int{1, 2, 3, 4, 5},
			Success: 7,
			Output:  []int{4, 0, 3},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			got := successfulPairs(test.Spells, test.Potions, test.Success)
			assert.DeepEqual(t, test.Output, got)
		})
	}
}
