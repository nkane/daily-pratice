package leetcode

import "fmt"

/*
	2352. Equal Row and Column Pairs
	Given a 0-Indexed n x n integer matrix grid, return the number of pairs (ri, ci) such that row ri and column
	ci are equal.
	A row and column pair is considered equal if they contain the same elements in the same order (i.e., an equal
	array).
*/

func equalPairs(grid [][]int) int {
	count := 0
	mRows := map[string]int{}
	for _, r := range grid {
		mRows[fmt.Sprintf("%v", r)]++
	}
	cols := make([]int, len(grid[0]))
	for c := 0; c < len(grid[0]); c++ {
		for r := 0; r < len(grid); r++ {
			cols[r] = grid[r][c]
		}
		key := fmt.Sprintf("%v", cols)
		if v, ok := mRows[key]; ok {
			count += v
		}
	}
	return count
}

const ARR_LENGTH = 200

func equalPairs_fast(grid [][]int) int {
	n := len(grid)

	rows := make(map[[ARR_LENGTH]int]int)
	cols := make(map[[ARR_LENGTH]int]int)

	count := 0

	for i := 0; i < n; i++ {
		row := [ARR_LENGTH]int{}
		col := [ARR_LENGTH]int{}

		for j := 0; j < n; j++ {
			row[j] = grid[i][j]
			col[j] = grid[j][i]
		}

		rows[row]++
		cols[col]++
	}

	for k, v := range rows {
		if cols[k] > 0 {
			count += cols[k] * v
		}
	}

	return count
}

type TrieNode struct {
	Count    int
	Children map[int]*TrieNode
}

type Trie struct {
	Root *TrieNode
}

func (t *Trie) Insert(a []int) {
	myTrie := t.Root
	for _, e := range a {
		if _, ok := myTrie.Children[e]; !ok {
			myTrie.Children[e] = &TrieNode{
				Count:    0,
				Children: make(map[int]*TrieNode),
			}
		}
		myTrie = myTrie.Children[e]
	}
	myTrie.Count++
}

func (t *Trie) Search(a []int) int {
	myTrie := t.Root
	for _, e := range a {
		if v, ok := myTrie.Children[e]; ok {
			myTrie = v
		} else {
			return 0
		}
	}
	return myTrie.Count
}

func equalPairs_trie(grid [][]int) int {
	t := Trie{
		Root: &TrieNode{
			Count:    0,
			Children: make(map[int]*TrieNode),
		},
	}
	count := 0
	n := len(grid)
	for _, row := range grid {
		t.Insert(row)
	}
	for c := 0; c < n; c++ {
		colArray := make([]int, n)
		for r := 0; r < n; r++ {
			colArray[r] = grid[r][c]
		}
		count += t.Search(colArray)
	}
	return count
}
