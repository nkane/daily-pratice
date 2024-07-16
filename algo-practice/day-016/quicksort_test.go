package practice

import (
	"testing"

	"gotest.tools/assert"
)

func TestQuicksort(t *testing.T) {
	a := []int{1024, 32, 2, 8, 64, 2048, 128, 0, 256}
	expected := []int{0, 2, 8, 32, 64, 128, 256, 1024, 2048}
	Quicksort(a, 0, len(a)-1)
	assert.DeepEqual(t, expected, a)
}
