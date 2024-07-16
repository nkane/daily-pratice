package leetcode

import (
	"testing"

	"gotest.tools/assert"
)

func TestMaxVowels(t *testing.T) {
	s := "abciiidef"
	k := 3
	got := maxVowels(s, k)
	assert.Assert(t, got == 3)

	s = "aeiou"
	k = 2
	got = maxVowels(s, k)
	assert.Assert(t, got == 2)

	s = "leetcode"
	k = 3
	got = maxVowels(s, k)
	assert.Assert(t, got == 2)

	s = "tryhard"
	k = 4
	got = maxVowels(s, k)
	assert.Assert(t, got == 1)

	s = "weallloveyou"
	k = 7
	got = maxVowels(s, k)
	assert.Assert(t, got == 4)
}

func TestMaxVowels_Solution(t *testing.T) {
	s := "abciiidef"
	k := 3
	got := maxVowels_solutions(s, k)
	assert.Assert(t, got == 3)

	s = "aeiou"
	k = 2
	got = maxVowels_solutions(s, k)
	assert.Assert(t, got == 2)

	s = "leetcode"
	k = 3
	got = maxVowels_solutions(s, k)
	assert.Assert(t, got == 2)

	s = "tryhard"
	k = 4
	got = maxVowels_solutions(s, k)
	assert.Assert(t, got == 1)

	s = "weallloveyou"
	k = 7
	got = maxVowels_solutions(s, k)
	assert.Assert(t, got == 4)
}
