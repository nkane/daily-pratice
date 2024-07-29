package leetcode

import (
	"testing"

	"gotest.tools/assert"
)

func TestDecodeString(t *testing.T) {
	s := "3[a]2[bc]"
	expected := "aaabcbc"
	got := decodeString(s)
	assert.Assert(t, expected == got)

	s = "3[a2[c]]"
	expected = "accaccacc"
	got = decodeString(s)
	assert.Assert(t, expected == got)

	s = "2[abc]3[cd]ef"
	expected = "abcabccdcdcdef"
	got = decodeString(s)
	assert.Assert(t, expected == got)
}
