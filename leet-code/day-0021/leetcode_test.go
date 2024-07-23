package leetcode

import (
	"testing"

	"gotest.tools/assert"
)

func TestCloseStrings(t *testing.T) {
	word1 := "abc"
	word2 := "bca"
	expected := true
	got := closeStrings(word1, word2)
	assert.Assert(t, expected == got)

	word1 = "a"
	word2 = "aa"
	expected = false
	got = closeStrings(word1, word2)
	assert.Assert(t, expected == got)

	word1 = "cabbba"
	word2 = "abbccc"
	expected = true
	got = closeStrings(word1, word2)
	assert.Assert(t, expected == got)
}

func TestCloseStrings_Bit(t *testing.T) {
	word1 := "abc"
	word2 := "bca"
	expected := true
	got := closeStrings_bit(word1, word2)
	assert.Assert(t, expected == got)

	word1 = "a"
	word2 = "aa"
	expected = false
	got = closeStrings_bit(word1, word2)
	assert.Assert(t, expected == got)

	word1 = "cabbba"
	word2 = "abbccc"
	expected = true
	got = closeStrings_bit(word1, word2)
	assert.Assert(t, expected == got)
}
