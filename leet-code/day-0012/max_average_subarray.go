package leetcode

import "math"

/*
643. Maximum Average Subarray I:
You are given an integer array `nums` consisting of `n` elements, and an integer `k`.
Find a contiguous subarray whose length is equal to `k` that has the maximum average
value and return this value. Any answer with a calculationg error les than 10^-5 will
be accepted.
*/

// NOTE(nick): sliding window approach
func findMaxAverage(nums []int, k int) float64 {
	total := 0
	for i := 0; i < k; i++ {
		total += nums[i]
	}
	maxTotal := total
	for i := 0; i < len(nums)-k; i++ {
		total -= nums[i]
		total += nums[i+k]
		if total > maxTotal {
			maxTotal = total
		}
	}
	return float64(maxTotal) / float64(k)
}

// NOTE(nick): cumulative sum
func findMaxAverageCumulative(nums []int, k int) float64 {
	sum := make([]int, len(nums))
	sum[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		sum[i] = sum[i-1] + nums[i]
	}
	result := float64(sum[k-1]) * 1.0 / float64(k)
	for i := k; i < len(nums); i++ {
		result = math.Max(result, float64(sum[i]-sum[i-k])*1.0/float64(k))
	}
	return result
}
