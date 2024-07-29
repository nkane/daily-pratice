package leetcode

import (
	"testing"

	"gotest.tools/assert"
)

func TestUniqueOccurrences(t *testing.T) {
	a := []int{1, 2, 2, 1, 1, 3}
	expected := true
	got := uniqueOccurrences(a)
	assert.Assert(t, expected == got)

	a = []int{1, 2}
	expected = false
	got = uniqueOccurrences(a)
	assert.Assert(t, expected == got)

	a = []int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}
	expected = true
	got = uniqueOccurrences(a)
	assert.Assert(t, expected == got)
}
