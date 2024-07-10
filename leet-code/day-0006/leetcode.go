package leetcode

import "math"

func increasingTriplet(nums []int) bool {
	firstNum := math.MaxInt
	secondNum := math.MaxInt
	for _, n := range nums {
		if n <= firstNum {
			firstNum = n
		} else if n <= secondNum {
			secondNum = n
		} else {
			return true
		}
	}
	return false
}
