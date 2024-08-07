package practice

import (
	"testing"

	"gotest.tools/assert"
)

func TestCountingSort(t *testing.T) {
	tests := []struct {
		A        []int
		Expected []int
	}{
		{
			A:        []int{128, 2, 64, 16, 8, 32, 4, 0, 256},
			Expected: []int{0, 2, 4, 8, 16, 32, 64, 128, 256},
		},
	}
	for _, test := range tests {
		got := CountingSort(test.A)
		assert.DeepEqual(t, test.Expected, got)
	}
}
