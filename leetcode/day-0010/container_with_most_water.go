package leetcode

/*
	11. Container With Most Water
You are given an integer array height of length n. There are n vertical lines drawn such that two endpoints
of the ith line are (i, 0) and (i, height[i]).

Find two lines that together with the x-axis form a container, such that the container contains the most
water.

Return the maximum amount of water a container can store.

Notice that you may not slant the container.

Input: height = [1, 8, 6, 2, 5, 4, 8, 3, 7]
Output: 49
*/

/*
func maxArea(height []int) int {
	lo := 0
	hi := 1
	result := 0
	for i := 0; i < len(height) && hi < len(height); i++ {
		if height[lo] < height[hi] {
			lo = hi
			hi++
		} else if height[lo] > height[hi] {
			tmp := (height[hi] * height[hi])
			if tmp > result {
				result = tmp
			}
			hi++
		} else if height[lo] == height[hi] {
			hi++
		}
	}
	return result
}
*/

func maxArea(height []int) int {
	max := 0
	left := 0
	right := len(height) - 1
	for left < right {
		min := 0
		if height[left] > height[right] {
			min = height[right]
		} else {
			min = height[left]
		}
		currentArea := min * (right - left)
		if currentArea > max {
			max = currentArea
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return max
}
