package practice

import (
	"testing"

	"gotest.tools/assert"
)

func TestQuicksort(t *testing.T) {
	a := []int{128, 2, 16, 0, 256, 32, 8, 64, 4}
	expected := []int{0, 2, 4, 8, 16, 32, 64, 128, 256}
	Quicksort(a, 0, len(a)-1)
	assert.DeepEqual(t, a, expected)
}
