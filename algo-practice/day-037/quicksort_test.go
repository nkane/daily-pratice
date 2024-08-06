package practice

import (
	"testing"

	"gotest.tools/assert"
)

func TestQuicksort(t *testing.T) {
	tests := []struct {
		A        []int
		Expected []int
	}{
		{
			A:        []int{128, 2, 64, 16, 8, 32, 256, 0, 4},
			Expected: []int{0, 2, 4, 8, 16, 32, 64, 128, 256},
		},
	}
	for _, test := range tests {
		Quicksort(test.A, 0, len(test.A)-1)
		assert.DeepEqual(t, test.Expected, test.A)
	}
}
