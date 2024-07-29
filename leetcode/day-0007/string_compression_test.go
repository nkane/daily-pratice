package leetcode

import (
	"testing"

	"gotest.tools/assert"
)

func TestCompress(t *testing.T) {
	chars := []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}
	got := compress(chars)
	assert.Assert(t, got == 6)

	chars = []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}
	got = leetcodeSolution(chars)
	assert.Assert(t, got == 6)

	chars = []byte{'a', 'b', 'c'}
	got = compress(chars)
	assert.Assert(t, got == 3)

	chars = []byte{'a', 'b', 'c'}
	got = leetcodeSolution(chars)
	assert.Assert(t, got == 3)
}
