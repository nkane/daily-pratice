package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGuessNumber(t *testing.T) {
	tests := []struct {
		Name   string
		N      int
		Output int
	}{
		{
			Name:   "Example 1",
			N:      10,
			Output: 6,
		},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			got := guessNumber(test.N)
			assert.Equal(t, test.Output, got)
		})
	}
}
