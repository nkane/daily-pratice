package leetcode

/*
	1137: N-th Tribonacci Number
*/

func tribonacci(n int) int {
	m := map[int]int{
		0: 0,
		1: 1,
		2: 1,
	}
	return dfs(n, m)
}

func dfs(i int, m map[int]int) int {
	if _, ok := m[i]; ok {
		return m[i]
	}
	answer := dfs(i-1, m) + dfs(i-2, m) + dfs(i-3, m)
	m[i] = answer
	return answer
}
