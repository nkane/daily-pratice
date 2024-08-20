package practice

import (
	"bytes"
	"os"
	"testing"

	"gotest.tools/assert"
)

func TestGraphDFS(t *testing.T) {
	g := Graph[int, int]{
		Visited:       make(map[int]struct{}),
		AdjacencyList: make(map[int][]int),
	}
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 2)
	g.AddEdge(2, 0)
	g.AddEdge(2, 3)
	g.AddEdge(3, 3)

	tests := []struct {
		Graph    *Graph[int, int]
		Expected string
	}{
		{
			Graph:    &g,
			Expected: "2 0 1 3 ",
		},
	}

	for _, test := range tests {
		oldOut := os.Stdout
		r, w, err := os.Pipe()
		assert.NilError(t, err)
		assert.NilError(t, err)
		os.Stdout = w

		g.DFS(2)

		w.Close()
		os.Stdout = oldOut
		buf := bytes.Buffer{}
		buf.ReadFrom(r)
		got := buf.String()
		assert.Assert(t, test.Expected == got, "\nexpected:\n%s, got:\n%s", test.Expected, got)
	}
}

/*
	leetcode 1926: Nearest Exit from Entrance in Maze
*/

func TestNearestExit(t *testing.T) {
	tests := []struct {
		Name    string
		Input   [][]byte
		Entrace []int
		Output  int
	}{
		{
			Name: "Example 1",
			Input: [][]byte{
				{'+', '+', '.', '+'},
				{'.', '.', '.', '+'},
				{'+', '+', '+', '.'},
			},
			Entrace: []int{1, 2},
			Output:  1,
		},
		{
			Name: "Example 2",
			Input: [][]byte{
				{'+', '+', '+'},
				{'.', '.', '.'},
				{'+', '+', '+'},
			},
			Entrace: []int{1, 0},
			Output:  2,
		},
		{
			Name: "Example 3",
			Input: [][]byte{
				{'.', '+'},
			},
			Entrace: []int{0, 0},
			Output:  -1,
		},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			got := nearestExit(test.Input, test.Entrace)
			assert.Assert(t, test.Output == got)
		})
	}
}

/*
	994: Rotting Oranges
	You are given a m x n grid where each cell can have one of three values:

	- 0, represents an empty cell
	- 1, represents a fresh orange
	- 2, represents a rotten organge

	Every minute, any fresh orange that is 4-directionally adjacent to a rotten orange becomes rotten.

	Return the minimum number of minutes that must elapse until no cell has a fresh orange. If this is
	impossible, return -1.
*/

func TestOrangesRotting(t *testing.T) {
	tests := []struct {
		Name   string
		Grid   [][]int
		Output int
	}{
		{
			Name: "Example 1",
			Grid: [][]int{
				{2, 1, 1},
				{1, 1, 0},
				{0, 1, 1},
			},
			Output: 4,
		},
		{
			Name: "Example 2",
			Grid: [][]int{
				{2, 1, 1},
				{0, 1, 1},
				{1, 0, 1},
			},
			Output: -1,
		},
		{
			Name: "Example 3",
			Grid: [][]int{
				{0, 2},
			},
			Output: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			got := orangesRotting(test.Grid)
			assert.Assert(t, test.Output == got, "expected: %d, got: %d", test.Output, got)
		})
	}
}

func TestOrangesRottingNeet(t *testing.T) {
	tests := []struct {
		Name   string
		Grid   [][]int
		Output int
	}{
		{
			Name: "Example 1",
			Grid: [][]int{
				{2, 1, 1},
				{1, 1, 0},
				{0, 1, 1},
			},
			Output: 4,
		},
		{
			Name: "Example 2",
			Grid: [][]int{
				{2, 1, 1},
				{0, 1, 1},
				{1, 0, 1},
			},
			Output: -1,
		},
		{
			Name: "Example 3",
			Grid: [][]int{
				{0, 2},
			},
			Output: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			got := orangesRottingNeet(test.Grid)
			assert.Assert(t, test.Output == got, "expected: %d, got: %d", test.Output, got)
		})
	}
}
