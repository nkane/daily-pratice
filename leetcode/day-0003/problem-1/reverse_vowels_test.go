package practice

import (
	"testing"

	"gotest.tools/assert"
)

func TestReverseVowels(t *testing.T) {
	tests := map[string]string{
		"hello":      "holle",
		"leetcode":   "leotcede",
		"race a car": "raca e car",
		"a.b,.":      "a.b,.",
		"Marge, let's \"went.\" I await news telegram.":             "Marge, let's \"went.\" i awaIt news telegram.",
		"No, it never propagates if I set a \"gap\" or prevention.": "No, it never propagates If i set a \"gap\" or prevention.",
	}
	for input, output := range tests {
		got := reverseVowels(input)
		assert.Assert(t, output == got, "expected: %s, got: %s", output, got)
	}
}
