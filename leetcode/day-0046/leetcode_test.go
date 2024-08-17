package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxScore(t *testing.T) {
	tests := []struct {
		Name   string
		Nums1  []int
		Nums2  []int
		K      int
		Output int64
	}{
		{
			Name: "Example 1",
			Nums1: []int{
				1, 3, 3, 2,
			},
			Nums2: []int{
				2, 1, 3, 4,
			},
			K:      3,
			Output: 12,
		},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			got := maxScore(test.Nums1, test.Nums2, test.K)
			assert.Equal(t, test.Output, got)
		})
	}
}
