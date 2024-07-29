package leetcode

import (
	"testing"
)

func TestRecentCounter(t *testing.T) {
	rc := Constructor()
	tests := []struct {
		input    int
		expected int
	}{
		{1, 1},
		{100, 2},
		{3001, 3},
		{3002, 3},
	}

	for _, test := range tests {
		result := rc.Ping(test.input)
		if result != test.expected {
			t.Errorf("For input %d, expected %d but got %d", test.input, test.expected, result)
		}
	}
}
