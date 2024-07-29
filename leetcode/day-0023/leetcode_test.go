package leetcode

import (
	"testing"

	"gotest.tools/assert"
)

func TestRemoveStars(t *testing.T) {
	s := "leet**cod*e"
	expected := "lecoe"
	got := removeStars(s)
	assert.Assert(t, expected == got)
}

func TestRemoveStars_two_pointers(t *testing.T) {
	s := "leet**cod*e"
	expected := "lecoe"
	got := removeStars_two_pointers(s)
	assert.Assert(t, expected == got)
}
