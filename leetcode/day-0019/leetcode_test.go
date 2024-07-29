package leetcode

import (
	"testing"

	"gotest.tools/assert"
)

func TestFindDifference(t *testing.T) {
	nums1 := []int{1, 2, 3}
	nums2 := []int{2, 4, 6}
	expected := [][]int{
		{1, 3},
		{4, 6},
	}
	got := findDifference(nums1, nums2)
	assertRowEqual(t, expected[0], got[0])
	assertRowEqual(t, expected[1], got[1])

	nums1 = []int{1, 2, 3, 3}
	nums2 = []int{1, 1, 2, 2}
	expected = [][]int{
		{3},
		{},
	}
	got = findDifference(nums1, nums2)
	assertRowEqual(t, expected[0], got[0])
	assertRowEqual(t, expected[1], got[1])
}

func assertRowEqual(t *testing.T, expected []int, got []int) {
	checked := len(expected)
	for _, e := range expected {
		for _, g := range got {
			if g == e {
				checked--
				break
			}
		}
	}
	assert.Assert(t, checked == 0)
}
