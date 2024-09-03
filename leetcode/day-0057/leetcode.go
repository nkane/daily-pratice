package leetcode

import (
	"math"
	"strconv"
	"strings"
)

/*
	62: Unique Paths
*/

func uniquePaths(m int, n int) int {
	d := make([][]int, m)
	for i := range d {
		d[i] = make([]int, n)
		for j := range d[i] {
			d[i][j] = 1
		}
	}
	for col := 1; col < m; col++ {
		for row := 1; row < n; row++ {
			d[col][row] = d[col-1][row] + d[col][row-1]
		}
	}
	return d[m-1][n-1]
}

/*
	1143: Longest Common Subsequence
*/

func longestCommonSubsequence(text1 string, text2 string) int {
	memo := make([][]int, len(text1)+1)
	for i := range memo {
		memo[i] = make([]int, len(text2)+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	return memoSolve(text1, text2, 0, 0, memo)
}

func memoSolve(text1, text2 string, p1, p2 int, memo [][]int) int {
	if p1 == len(text1) || p2 == len(text2) {
		return 0
	}
	if memo[p1][p2] != -1 {
		return memo[p1][p2]
	}

	// Option 1: Skip the current character in text1
	option1 := memoSolve(text1, text2, p1+1, p2, memo)

	// Option 2: Include the current character from text1 if it matches with text2
	option2 := 0
	if p1 < len(text1) {
		if firstOccurrence := strings.Index(text2[p2:], string(text1[p1])); firstOccurrence != -1 {
			option2 = 1 + memoSolve(text1, text2, p1+1, p2+firstOccurrence+1, memo)
		}
	}

	// Store the result in memo and return
	memo[p1][p2] = int(math.Max(float64(option1), float64(option2)))
	return memo[p1][p2]
}

/*
	714: Best Time to Buy and Sell Stock with Transaction Fee
*/

func maxProfit(prices []int, fee int) int {
	cache := make(map[string]int)
	return findProfit(0, fee, 0, prices, cache)
}

func findProfit(idx, fee, state int, prices []int, cache map[string]int) int {
	if idx == len(prices) {
		return 0
	}

	key := strconv.Itoa(idx) + "-" + strconv.Itoa(state)
	if value, ok := cache[key]; ok {
		return value
	}

	doNothing := findProfit(idx+1, fee, state, prices, cache)
	doSomething := 0

	if state == 1 { // Bought
		doSomething = findProfit(idx+1, fee, 0, prices, cache) + prices[idx] - fee
	} else {
		doSomething = findProfit(idx+1, fee, 1, prices, cache) - prices[idx]
	}

	res := max(doNothing, doSomething)
	cache[key] = res
	return res
}

/*
	72: Edit Distance
*/

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := range m + 1 {
		dp[i] = make([]int, n+1)
		dp[i][0] = i
	}
	for j := range n + 1 {
		dp[0][j] = j
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if word1[i] == word2[j] {
				dp[i+1][j+1] = dp[i][j]
			} else {
				dp[i+1][j+1] = 1 + Min(dp[i][j], Min(dp[i+1][j], dp[i][j+1]))
			}
		}
	}
	return dp[m][n]
}

func Min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

/*
	338: Counting Bits
*/

func countBits(n int) []int {
	ones := make([]int, n+1)
	for x := 0; x <= n; x++ {
		ones[x] = ones[x>>1] + x&1
	}
	return ones
}

/*
	136: Single Number
*/

func singleNumber(nums []int) int {
	a := 0
	for _, i := range nums {
		a ^= i
	}
	return a
}

/*
	1318: Minimum Flips to Make a OR b Equal to c
*/

func minFlips(a int, b int, c int) int {
	ans := 0
	for a > 0 || b > 0 || c > 0 {
		ai, bi, ci := a%2, b%2, c%2
		a, b, c = a/2, b/2, c/2
		if ai|bi != ci {
			ans++
			if ci == 0 && ai == 1 && bi == 1 {
				ans++
			}
		}
	}
	return ans
}
