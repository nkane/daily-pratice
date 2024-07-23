package practice

import (
	"testing"

	"gotest.tools/assert"
)

func TestCountingSort(t *testing.T) {
	a := []int{128, 2, 4, 32, 16, 8, 0, 64}
	expected := []int{0, 2, 4, 8, 16, 32, 64, 128}
	got := CountingSort(a)
	assert.DeepEqual(t, expected, got)
}
