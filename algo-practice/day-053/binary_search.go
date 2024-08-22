package practice

import "golang.org/x/exp/constraints"

func BinarySearch[T constraints.Ordered](a []T, target T) (int, int) {
	count := 0
	lo := 0
	hi := len(a) - 1
	for lo <= hi {
		count++
		mid := lo + (hi-lo)/2
		guess := a[mid]
		if guess == target {
			return mid, count
		}
		if guess > target {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return -1, count
}

/*
162: Find Peak Element
*/
func findPeakElement(nums []int) int {
	lo := 0
	hi := len(nums) - 1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if mid > 0 && nums[mid] < nums[mid-1] {
			hi = mid - 1
		} else if mid < len(nums)-1 && nums[mid] < nums[mid+1] {
			lo = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
