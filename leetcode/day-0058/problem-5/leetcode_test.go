package leetcode

import (
	"testing"

	"gotest.tools/assert"
)

func TestDailyTemperatures(t *testing.T) {
	tests := []struct {
		Name   string
		Temps  []int
		Output []int
	}{
		{
			Name: "Example 1",
			Temps: []int{
				73, 74, 75, 71, 69, 72, 76, 73,
			},
			Output: []int{
				1, 1, 4, 2, 1, 1, 0, 0,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			got := dailyTemperatures(test.Temps)
			assert.DeepEqual(t, test.Output, got)
		})
	}
}
