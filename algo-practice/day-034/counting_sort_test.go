package practice

import (
	"testing"

	"gotest.tools/assert"
)

func TestCountingSort(t *testing.T) {
	a := []int{128, 2, 256, 64, 0, 16, 4, 8, 32}
	expected := []int{0, 2, 4, 8, 16, 32, 64, 128, 256}
	got := CountingSort(a)
	assert.DeepEqual(t, expected, got)
}
