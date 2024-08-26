package leetcode

/*
	746: Min Cost Climbing Stairs
*/

func minCostClimbingStairs(cost []int) int {
	memo := make(map[int]int)
	return minimumCost(len(cost), cost, memo)
}

func minimumCost(i int, cost []int, memo map[int]int) int {
	if i <= 1 {
		return 0
	}
	if val, ok := memo[i]; ok {
		return val
	}
	downOne := cost[i-1] + minimumCost(i-1, cost, memo)
	downTwo := cost[i-2] + minimumCost(i-2, cost, memo)
	memo[i] = min(downOne, downTwo)
	return memo[i]
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
