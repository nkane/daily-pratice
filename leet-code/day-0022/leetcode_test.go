package leetcode

import (
	"testing"

	"gotest.tools/assert"
)

func TestEqualPairs(t *testing.T) {
	grid := [][]int{
		{3, 2, 1},
		{1, 7, 6},
		{2, 7, 7},
	}
	expected := 1
	got := equalPairs(grid)
	assert.Assert(t, expected == got, "expected: %d, got: %d\n", expected, got)

	grid = [][]int{
		{3, 1, 2, 2},
		{1, 4, 4, 5},
		{2, 4, 2, 2},
		{2, 4, 2, 2},
	}
	expected = 3
	got = equalPairs(grid)
	assert.Assert(t, expected == got, "expected: %d, got: %d\n", expected, got)
}

func TestEqualPairs_Trie(t *testing.T) {
	grid := [][]int{
		{3, 2, 1},
		{1, 7, 6},
		{2, 7, 7},
	}
	expected := 1
	got := equalPairs_trie(grid)
	assert.Assert(t, expected == got, "expected: %d, got: %d\n", expected, got)
}
