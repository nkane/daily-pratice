package leetcode

/*
	485: Max Consecutive Ones
	Given a binary array nums, return the maximum number of consecutive 1's in the array.
*/

func findMaxConsecutiveOnes(nums []int) int {
	max := 0
	temp := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			temp++
			if temp > max {
				max = temp
			}
		} else {
			temp = 0
		}
	}
	return max
}
