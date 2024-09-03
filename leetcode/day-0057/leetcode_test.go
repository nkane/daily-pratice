package leetcode

import (
	"testing"

	"gotest.tools/assert"
)

func TestUniquePaths(t *testing.T) {
	tests := []struct {
		Name   string
		M      int
		N      int
		Output int
	}{
		{
			Name:   "Example 1",
			M:      3,
			N:      7,
			Output: 28,
		},
		{
			Name:   "Example 2",
			M:      3,
			N:      2,
			Output: 3,
		},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			got := uniquePaths(test.M, test.N)
			assert.Assert(t, test.Output == got)
		})
	}
}

func TestLongestCommonSubsequence(t *testing.T) {
	tests := []struct {
		Name   string
		Text1  string
		Text2  string
		Output int
	}{
		{
			Name:   "Example 1",
			Text1:  "abcde",
			Text2:  "ace",
			Output: 3,
		},
		{
			Name:   "Example 2",
			Text1:  "abc",
			Text2:  "abc",
			Output: 3,
		},
		{
			Name:   "Example 3",
			Text1:  "abc",
			Text2:  "def",
			Output: 0,
		},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			got := longestCommonSubsequence(test.Text1, test.Text2)
			assert.Assert(t, test.Output == got)
		})
	}
}

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		Name   string
		Prices []int
		Fee    int
		Output int
	}{
		{
			Name:   "Example 1",
			Prices: []int{1, 3, 2, 8, 4, 9},
			Fee:    2,
			Output: 8,
		},
		{
			Name:   "Example 2",
			Prices: []int{1, 3, 7, 5, 10, 3},
			Fee:    3,
			Output: 6,
		},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			got := maxProfit(test.Prices, test.Fee)
			assert.Assert(t, test.Output == got)
		})
	}
}

func TestMinDistance(t *testing.T) {
	tests := []struct {
		Name   string
		Word1  string
		Word2  string
		Output int
	}{
		{
			Name:   "Example 1",
			Word1:  "horse",
			Word2:  "ros",
			Output: 3,
		},
		{
			Name:   "Example 2",
			Word1:  "intention",
			Word2:  "execution",
			Output: 5,
		},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			got := minDistance(test.Word1, test.Word2)
			assert.Assert(t, test.Output == got)
		})
	}
}

func TestCountingBits(t *testing.T) {
	tests := []struct {
		Name   string
		N      int
		Output []int
	}{
		{
			Name:   "Example 1",
			N:      2,
			Output: []int{0, 1, 1},
		},
		{
			Name:   "Example 2",
			N:      5,
			Output: []int{0, 1, 1, 2, 1, 2},
		},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			got := countBits(test.N)
			assert.DeepEqual(t, test.Output, got)
		})
	}
}

func TestSingleNumber(t *testing.T) {
	tests := []struct {
		Name   string
		Nums   []int
		Output int
	}{
		{
			Name:   "Example 1",
			Nums:   []int{2, 2, 1},
			Output: 1,
		},
		{
			Name:   "Example 2",
			Nums:   []int{4, 1, 2, 1, 2},
			Output: 4,
		},
		{
			Name:   "Example 3",
			Nums:   []int{1},
			Output: 1,
		},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			got := singleNumber(test.Nums)
			assert.Assert(t, test.Output == got)
		})
	}
}

func TestMinFlips(t *testing.T) {
	tests := []struct {
		Name   string
		A      int
		B      int
		C      int
		Output int
	}{
		{
			Name:   "Example 1",
			A:      2,
			B:      6,
			C:      5,
			Output: 3,
		},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			got := minFlips(test.A, test.B, test.C)
			assert.Assert(t, test.Output == got)
		})
	}
}
