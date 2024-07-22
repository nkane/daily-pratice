package practice

import (
	"testing"

	"gotest.tools/assert"
)

func TestCountingSort(t *testing.T) {
	a := []int{128, 2, 16, 64, 8, 0, 32}
	expected := []int{0, 2, 8, 16, 32, 64, 128}
	got := CountingSort(a)
	assert.DeepEqual(t, expected, got)
}
