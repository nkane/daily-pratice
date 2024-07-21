package practice

import (
	"testing"

	"gotest.tools/assert"
)

func TestCountingSort(t *testing.T) {
	a := []int{1024, 8, 16, 32, 0, 4, 128, 256, 64, 2}
	expected := []int{0, 2, 4, 8, 16, 32, 64, 128, 256, 1024}
	got := CountingSort(a)
	assert.DeepEqual(t, expected, got)
}
