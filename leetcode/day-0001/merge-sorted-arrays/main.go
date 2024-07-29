package main

import (
	"log"
	"reflect"
	"sort"
)

/*
88. Merge Sorted Array

You are given two integer arrays nums1 and nums2, sorted in non-decreasing order, and two integers m and n, representing the number of elements in nums1 and nums2 respectively.
Merge nums1 and nums2 into a single array sorted in non-decreasing order.
The final sorted array should not be returned by the function, but instead be stored inside the array nums1. To accommodate this, nums1 has a length of m + n, where the first m elements denote the elements that should be merged, and the last n elements are set to 0 and should be ignored. nums2 has a length of n.

Example 1:

Input: nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
Output: [1,2,2,3,5,6]
Explanation: The arrays we are merging are [1,2,3] and [2,5,6].
The result of the merge is [1,2,2,3,5,6] with the underlined elements coming from nums1.
Example 2:

Input: nums1 = [1], m = 1, nums2 = [], n = 0
Output: [1]
Explanation: The arrays we are merging are [1] and [].
The result of the merge is [1].
Example 3:

Input: nums1 = [0], m = 0, nums2 = [1], n = 1
Output: [1]
Explanation: The arrays we are merging are [] and [1].
The result of the merge is [1].
Note that because m = 0, there are no elements in nums1. The 0 is only there to ensure the merge result can fit in nums1.
*/

func main() {
	// nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3
	// expected: [1,2,2,3,5,6], pass
	merge(nums1, m, nums2, n)
	if !reflect.DeepEqual(nums1, []int{1, 2, 2, 3, 5, 6}) {
		log.Fatalf("failed 1st test case")
	}

	nums1 = []int{1}
	m = 1
	nums2 = []int{}
	n = 0
	merge(nums1, m, nums2, n)
	// expected: [1], pass
	if !reflect.DeepEqual(nums1, []int{1}) {
		log.Fatalf("failed 2nd test case")
	}

	nums1 = []int{0}
	m = 0
	nums2 = []int{1}
	n = 1
	merge(nums1, m, nums2, n)
	// expected: [1], pass
	if !reflect.DeepEqual(nums1, []int{1}) {
		log.Fatalf("failed 2nd test case")
	}
	// incorrect answer: [1,2,3,6,5,4]
	/*
		nums1 = [4,5,6,0,0,0]
		m = 3
		nums2 = [1,2,3]
		n = 3
	*/
	nums1 = []int{4, 5, 6, 0, 0, 0}
	m = 3
	nums2 = []int{1, 2, 3}
	n = 3
	merge(nums1, m, nums2, n)
	if !reflect.DeepEqual(nums1, []int{1, 2, 3, 4, 5, 6}) {
		log.Fatalf("failed 2nd test case")
	}
}

func mergeFAILED(nums1 []int, m int, nums2 []int, n int) {
	for j := 0; j < n; j++ {
		for i := 0; i < m; i++ {
			if nums1[i] > nums2[j] {
				// swap
				nums1[i], nums2[j] = nums2[j], nums1[i]
			}
		}
	}
	idx := m
	if n > 0 {
		if nums2[0] > nums2[n-1] {
			// reverse order
			for i := len(nums2) - 1; i >= 0; i-- {
				nums1[idx] = nums2[i]
				idx++
			}
		} else {
			// normal order
			for i := 0; i < len(nums2); i++ {
				nums1[idx] = nums2[i]
				idx++
			}
		}
	}
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	j := 0
	for i := m; i < (m + n); i++ {
		nums1[i] = nums2[j]
		j++
	}
	sort.Ints(nums1)
}

func mergeBetterSolution(nums1 []int, m int, nums2 []int, n int) {
	for n != 0 {
		if m != 0 && nums1[m-1] > nums2[n-1] {
			nums1[m+n-1] = nums1[m-1]
			m--
		} else {
			nums1[m+n-1] = nums2[n-1]
			n--
		}
	}
}
