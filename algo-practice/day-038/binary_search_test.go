package practice

import (
	"testing"

	"gotest.tools/assert"
)

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		A             []int
		Target        int
		ExpectedIdx   int
		ExpectedCount int
	}{
		{
			A:             []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
			Target:        8,
			ExpectedIdx:   7,
			ExpectedCount: 1,
		},
	}
	for _, test := range tests {
		idx, count := BinarySearch(test.A, test.Target)
		assert.Assert(t, idx == test.ExpectedIdx)
		assert.Assert(t, count == test.ExpectedCount)
	}
}
