package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTotalCost(t *testing.T) {
	tests := []struct {
		Name       string
		Costs      []int
		K          int
		Candidates int
		Output     int64
	}{
		{
			Name: "Example 1",
			Costs: []int{
				17, 12, 10, 2, 7, 2, 11, 20, 8,
			},
			K:          3,
			Candidates: 4,
			Output:     11,
		},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			got := totalCost(test.Costs, test.K, test.Candidates)
			assert.Equal(t, test.Output, got)
		})
	}
}
