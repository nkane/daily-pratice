package leetcode

import "sort"

/*
	452: Minimum Number of Arrows to Burst Balloons
*/

func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})
	i, ans := 0, 0
	for i < len(points) {
		x := points[i][1]
		ans++
		i++
		for i < len(points) && x >= points[i][0] {
			i++
		}
	}
	return ans
}
