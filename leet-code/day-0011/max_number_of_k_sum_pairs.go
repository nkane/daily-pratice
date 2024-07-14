package leetcode

/*
1679. Max Numbers of K-Sum Pairs

You are given an integer array `nums` and an integer `k`.

In one operation, you can pick two numbers from the array whose sum equals `k` and remove them from the
array.

Return the minimum number of operations you can perform on the array.

Example 1:
Input: nums = [1, 2, 3, 4, 5], k = 5
Output: 2
Explaination: Start with nums = [1, 2, 3, 4, 5]:
- Remove numbers 1 and 4, then nums = [2, 3]
- Remove numbers 2 and 3, then nums = []
There are no more pairs that sum up to 5, hence the total of 2 operations

Example 2:
Input: nums = [3, 1, 3, 4, 3], k = 6
Output: 1
Explaination: Start with nums = [3, 1, 3, 4, 3]:
- Remove the first two 3's, them nums = [1, 4, 3]
There are no more pairs that sum up to 6, hence the total of 1 operations
*/

func maxOperations(nums []int, k int) int {
	ops := 0
	m := map[int]int{}
	for _, n := range nums {
		m[n]++
	}
	for i := 0; i < len(nums); i++ {
		current := nums[i]
		complement := k - nums[i]
		if m[current] > 0 && m[complement] > 0 {
			if (current == complement) && m[current] < 2 {
				continue
			}
			m[current]--
			m[complement]--
			ops++
		}
	}
	return ops
}
