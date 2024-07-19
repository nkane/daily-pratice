package practice

import (
	"testing"

	"gotest.tools/assert"
)

func TestQuicksort(t *testing.T) {
	a := []int{128, 2, 256, 64, 8, 4, 0, 1024}
	expected := []int{0, 2, 4, 8, 64, 128, 256, 1024}
	Quicksort(a, 0, len(a)-1)
	assert.DeepEqual(t, expected, a)
}
