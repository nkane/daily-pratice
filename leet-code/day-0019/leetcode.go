package leetcode

/*
	2215. Find the Difference of Two Arrays

	Given two 0-indexed integer arrays nums1 and nums2, return a list answer of size 2 where:
	- answer[0] is a list of all distinct integers in nums1 which are not present in nums2.
	- answer[1] is a list of all distinct integers in nums2 which are not present in nums1....
	Note that the integers in the list may be returned in any order.
*/

func findDifference(nums1 []int, nums2 []int) [][]int {
	m1 := map[int]struct{}{}
	for _, e := range nums1 {
		m1[e] = struct{}{}
	}
	m2 := map[int]struct{}{}
	for _, e := range nums2 {
		m2[e] = struct{}{}
	}
	output := make([][]int, 2)
	output[0] = make([]int, 0)
	output[1] = make([]int, 0)
	for k := range m1 {
		if _, ok := m2[k]; !ok {
			output[0] = append(output[0], k)
		}
	}
	for k := range m2 {
		if _, ok := m1[k]; !ok {
			output[1] = append(output[1], k)
		}
	}
	return output
}
