package leetcode

/*
	1207. Unique Number of Occurrences
	Given an array of integers arr, return true if the number of occurences of each balue in the
	array is unique or false otherwise.
*/

func uniqueOccurrences(arr []int) bool {
	m := map[int]int{}
	for _, e := range arr {
		m[e]++
	}
	counts := map[int]struct{}{}
	for _, v := range m {
		if _, ok := counts[v]; ok {
			return false
		}
		counts[v] = struct{}{}
	}
	return true
}
