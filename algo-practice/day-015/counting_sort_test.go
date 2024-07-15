package practice

import (
	"testing"

	"gotest.tools/assert"
)

func TestCountingSort(t *testing.T) {
	a := []int{3, 6, 1, 5, 2, 4, 8}
	expected := []int{1, 2, 3, 4, 5, 6, 8}
	got := CountingSort(a)
	assert.DeepEqual(t, expected, got)

	a = []int{1024, 256, 16, 0, 32, 2, 2048, 128, 64}
	expected = []int{0, 2, 16, 32, 64, 128, 256, 1024, 2048}
	got = CountingSort(a)
	assert.DeepEqual(t, expected, got)
}
