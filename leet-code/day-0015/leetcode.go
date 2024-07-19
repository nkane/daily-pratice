package leetcode

/*
Given a binary array nums, you should delete one element from it.

Return the size of the longest non-empty subarray containing only 1's in the resulting array.
Return 0 if there is no such subarray
*/

func longestSubarray(nums []int) int {
	head := 0
	max := 0
	n := 0
	m := 0
	allOne := true
	for head < len(nums) {
		if nums[head] == 1 {
			n++
			if head == len(nums)-1 {
				c := m + n
				if c > max {
					max = c
				}
			}
		} else if nums[head] == 0 {
			allOne = false
			c := m + n
			if c > max {
				max = c
			}
			m = n
			n = 0
		}
		head++
	}
	if allOne {
		max--
	}
	return max
}

func longestSubarray_Solution(nums []int) int {
	// Number of zero's in the window.
	zeroCount := 0
	longestWindow := 0
	// Left end of the window.
	start := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			zeroCount++
		}

		// Shrink the window until the zero counts come under the limit.
		for zeroCount > 1 {
			if nums[start] == 0 {
				zeroCount--
			}
			start++
		}

		longestWindow = max(longestWindow, i-start)
	}

	return longestWindow
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
