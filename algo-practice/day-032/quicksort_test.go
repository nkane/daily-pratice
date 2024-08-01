package practice

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestQuicksort(t *testing.T) {
	a := []int{128, 4, 32, 8, 2, 0, 16, 64, 256}
	expected := []int{0, 2, 4, 8, 16, 32, 64, 128, 256}
	Quicksort(a, 0, len(a)-1)
	assert.DeepEqual(t, expected, a)
}
