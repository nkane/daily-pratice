package leetcode

/*
	724. Find Pivot Index

	Given an array of integers `nums`, calculate the pivot index of this array.

	The pivot idnex is the index where the sum of all the numbers are strictly to the left
	of the index is equal to the sum of all the numbers strictly to the index's right.

	If the index is on the left edge of this array, then the left sum is 0 because there are no
	elements to the left. This also applies to the right edge of the array.

	Return the leftmost pivot index, if no such index exists, return -1.
*/

// NOTE(nick): failed attempt
/*
func pivotIndex(nums []int) int {
	leftSum := make([]int, len(nums))
	leftSum[0] = nums[0]
	rightSum := make([]int, len(nums))
	rightSum[0] = nums[len(nums)-1]
	for i := 1; i < len(nums); i++ {
		if leftSum[i-1] == rightSum[i-1] {
			return i
		}
		leftSum[i] = (nums[i] + leftSum[i-1])
		j := (len(nums) - 1) - i
		rightSum[i] = (nums[j]) + rightSum[i-1]
	}
	return -1
}
*/

// NOTE(nick): editorial solution
func pivotIndex(nums []int) int {
	totalSum := 0
	for _, num := range nums {
		totalSum += num
	}
	leftSum := 0
	for i, num := range nums {
		// totalSum - leftSum - num is the right sum at index i
		if leftSum == (totalSum - leftSum - num) {
			return i
		}
		leftSum += num
	}
	return -1
}
