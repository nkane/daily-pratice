package reverse_words

import (
	"testing"

	"gotest.tools/assert"
)

func TestReverseWord(t *testing.T) {
	tests := []struct {
		Input    string
		Expected string
	}{
		{
			Input:    "the sky is blue",
			Expected: "blue is sky the",
		},
		{
			Input:    "  hello world  ",
			Expected: "world hello",
		},
		{
			Input:    "a good   example",
			Expected: "example good a",
		},
	}
	for _, test := range tests {
		got := reverseWords(test.Input)
		assert.Assert(t, test.Expected == got, "reverseWord expected: %s, got: %s", test.Expected, got)

		got = reverseWordsFast(test.Input)
		assert.Assert(t, test.Expected == got, "reverseWordFast expected: %s, got: %s", test.Expected, got)
	}
}
