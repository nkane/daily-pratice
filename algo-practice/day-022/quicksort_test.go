package practice

import (
	"testing"

	"gotest.tools/assert"
)

func TestQuicksort(t *testing.T) {
	a := []int{128, 2, 16, 64, 8, 0, 32}
	expected := []int{0, 2, 8, 16, 32, 64, 128}
	Quicksort(a, 0, len(a)-1)
	assert.DeepEqual(t, expected, a)
}
