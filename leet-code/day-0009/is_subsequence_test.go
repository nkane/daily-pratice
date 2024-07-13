package leetcode

import (
	"testing"

	"gotest.tools/assert"
)

func TestIsSubsequence(t *testing.T) {
	sv := "abc"
	tv := "ahbgdc"
	expected := true
	got := isSubsequence(sv, tv)
	assert.Assert(t, expected == got)

	sv = "axc"
	tv = "ahbgdc"
	expected = false
	got = isSubsequence(sv, tv)
	assert.Assert(t, expected == got)

	sv = ""
	tv = "ahbgdc"
	expected = false
	got = isSubsequence(sv, tv)
	assert.Assert(t, expected == got)
}
