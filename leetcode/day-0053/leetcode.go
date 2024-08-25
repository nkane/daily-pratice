package leetcode

/*
	216: Combination Sum III
*/

func backtrack(remain, k int, comb []int, nextStart int, results *[][]int) {
	if remain == 0 && len(comb) == k {
		// Make a deep copy of comb and append it to results
		combCopy := make([]int, len(comb))
		copy(combCopy, comb)
		*results = append(*results, combCopy)
		return
	} else if remain < 0 || len(comb) == k {
		// Exceeded the scope, no need to explore further
		return
	}

	// Iterate through the reduced list of candidates
	for i := nextStart; i < 9; i++ {
		comb = append(comb, i+1)
		backtrack(remain-i-1, k, comb, i+1, results)
		comb = comb[:len(comb)-1]
	}
}

func combinationSum3(k int, n int) [][]int {
	var results [][]int
	var comb []int
	backtrack(n, k, comb, 0, &results)
	return results
}
