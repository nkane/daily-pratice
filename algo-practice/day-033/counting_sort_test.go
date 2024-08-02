package practice

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestCountingSort(t *testing.T) {
	a := []int{128, 0, 2, 32, 64, 8, 256, 4, 16}
	expected := []int{0, 2, 4, 8, 16, 32, 64, 128, 256}
	got := CountingSort(a)
	assert.DeepEqual(t, expected, got)
}
