package practice

import (
	"testing"

	"gotest.tools/assert"
)

func TestCountingSort(t *testing.T) {
	a := []int{6, 2, 5, 3, 1, 0, 4}
	expected := []int{0, 1, 2, 3, 4, 5, 6}
	got := CountingSort(a)
	assert.DeepEqual(t, expected, got)
}
