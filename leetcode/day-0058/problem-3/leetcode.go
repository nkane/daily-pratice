package leetcode

import "sort"

/*
	435: Non-overlapping Intervals
*/
// Time: O(nlogn + n) = O(nlogn)
// Space: O(1)
func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) <= 1 {
		return 0
	}

	// Time: O(nlogn)
	sortIntervals(intervals)

	// Time: O(n)
	nbNonOverlappingIntervals := numberOfNonOverlappingIntervals(intervals)

	return len(intervals) - nbNonOverlappingIntervals
}

// Time: O(n)
// Space: O(1)
func numberOfNonOverlappingIntervals(intervals [][]int) int {
	currentInterval := intervals[0]
	var nbNonOverlappingIntervals int = 1

	// Time: O(n)
	for _, interval := range intervals[1:] {
		if interval[0] >= currentInterval[1] {
			currentInterval = interval
			nbNonOverlappingIntervals++
		}
	}

	return nbNonOverlappingIntervals
}

// Time: O(nlogn)
// Space: O(1)
func sortIntervals(intervals [][]int) {
	// Time: O(nlogn)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
}
