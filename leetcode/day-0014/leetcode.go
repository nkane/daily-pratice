package leetcode

/*
	1004. Max Consecutive Ones III
	Given a binary array `nums` and an inter `k`, return the maximum number of consecutive 1's in an array if you
	can flip at most k 0's

	Example 1:
	Input: nums = [1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0], k = 2
	Output: 6
	Explaination: [1, 1, 1, 0, 0, _1_, 1, 1, 1, 1, _1_]
	Bolded numbers were flipped from 0 to 1, the longest subarray is underlined.

	Example 2:
	Input: nums = [0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1], k = 3
	Output: 10
	Explanation: [0, 0, 1, 1, _1_, _1_, 1, 1, 1, _1_, 1, 1, 0, 0, 0, 1, 1, 1, 1]
	Bolded numbers were flipped from 0 to 1. The longest subarray is underlined
*/

func longestOne(nums []int, k int) int {
	left := 0
	right := 0
	n := len(nums)
	for right = 0; right < n; right++ {
		if nums[right] == 0 {
			k--
		}
		if k < 0 {
			k += 1 - nums[left]
			left++
		}
	}
	return right - left
}
