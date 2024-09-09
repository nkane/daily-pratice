package leetcode

import "math"

/*
	1295: Find Numbers with Even Numbers of Digits
	Given an array nums of integers, return how many of them contain an even number of digits.
*/

func findNumbers(nums []int) int {
	res := 0
	for _, num := range nums {
		digits := 0
		for num > 0 {
			digits++
			num /= 10
		}
		if (digits & 0x01) == 0 {
			res++
		}
	}
	return res
}

/*
	977: Squares of a Sorted Array
	Given an integer array nums sorted in non-decreasing order, return an array of the squares of each number sorted
	in non-decreasing order.
*/

func sortedSquares(nums []int) []int {
	n := len(nums)
	result := make([]int, n)
	left, right := 0, n-1
	for i := n - 1; i >= 0; i-- {
		var square int
		if math.Abs(float64(nums[left])) < math.Abs(float64(nums[right])) {
			square = nums[right]
			right--
		} else {
			square = nums[left]
			left++
		}
		result[i] = square * square
	}
	return result
}
