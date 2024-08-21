package leetcode

/*
	162: Find Peak Element
*/

func findPeakElement(nums []int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		m := l + (r-l)/2
		if m > 0 && nums[m] < nums[m-1] {
			// left neighbor greater
			r = m - 1
		} else if m < len(nums)-1 && nums[m] < nums[m+1] {
			// right neighbor greater
			l = m + 1
		} else {
			return m
		}
	}
	return -1
}
