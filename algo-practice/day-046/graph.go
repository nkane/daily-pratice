package practice

import (
	"fmt"
)

type Graph[K comparable, T any] struct {
	Visited       map[K]struct{}
	AdjacencyList map[K][]K
}

func (g *Graph[K, T]) AddEdge(v K, w K) {
	g.AdjacencyList[v] = append(g.AdjacencyList[v], w)
}

func (g *Graph[K, T]) DFS(v K) {
	g.Visited[v] = struct{}{}
	fmt.Print(v, " ")
	for _, edge := range g.AdjacencyList[v] {
		if _, visited := g.Visited[edge]; !visited {
			g.DFS(edge)
		}
	}
}

type Cell struct {
	X    int
	Y    int
	Step int
}

func nearestExit(maze [][]byte, entrance []int) int {
	rows, cols := len(maze), len(maze[0])
	dirs := [][]int{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}
	startRow, startCol := entrance[0], entrance[1]
	maze[startRow][startCol] = '+'
	queue := [][]int{
		{startRow, startCol, 0},
	}
	for len(queue) > 0 {
		currRow, currCol, currDistance := queue[0][0], queue[0][1], queue[0][2]
		queue = queue[1:]
		for _, d := range dirs {
			nextRow, nextCol := currRow+d[0], currCol+d[1]
			// boundary and available slot check
			if nextRow >= 0 && nextRow < rows &&
				nextCol >= 0 && nextCol < cols &&
				maze[nextRow][nextCol] == '.' {
				// reached an exit
				if nextRow == 0 || nextRow == rows-1 ||
					nextCol == 0 || nextCol == cols-1 {
					return currDistance + 1
				}
				maze[nextRow][nextCol] = '+'
				queue = append(queue, []int{nextRow, nextCol, currDistance + 1})
			}
		}
	}
	return -1
}

func orangesRotting(grid [][]int) int {
	queue := [][]int{}
	freshOranges := 0
	rows := len(grid)
	cols := len(grid[0])
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == 2 {
				queue = append(queue, []int{r, c})
			} else if grid[r][c] == 1 {
				freshOranges++
			}
		}
	}
	queue = append(queue, []int{-1, -1})
	minutesElapsed := -1
	directions := [][]int{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}
	for len(queue) > 0 {
		row, col := queue[0][0], queue[0][1]
		queue = queue[1:]
		if row == -1 {
			minutesElapsed++
			if len(queue) > 0 {
				queue = append(queue, []int{-1, -1})
			}
		} else {
			for _, d := range directions {
				nRow, nCol := row+d[0], col+d[1]
				if rows > nRow && nRow >= 0 &&
					cols > nCol && nCol >= 0 {
					if grid[nRow][nCol] == 1 {
						grid[nRow][nCol] = 2
						freshOranges--
						queue = append(queue, []int{nRow, nCol})
					}
				}
			}
		}
	}
	if freshOranges == 0 {
		return minutesElapsed
	}
	return -1
}

func orangesRottingNeet(grid [][]int) int {
	q := [][]int{}
	time, fresh := 0, 0
	rows, cols := len(grid), len(grid[0])
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == 1 {
				fresh++
			} else if grid[r][c] == 2 {
				q = append(q, []int{r, c})
			}
		}
	}
	directions := [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}
	for len(q) > 0 && fresh > 0 {
		l := len(q)
		for i := 0; i < l; i++ {
			r, c := q[0][0], q[0][1]
			q = q[1:]
			for _, d := range directions {
				row, col := d[0]+r, d[1]+c
				if row < 0 || row == len(grid) ||
					col < 0 || col == len(grid[0]) ||
					grid[row][col] != 1 {
					continue
				}
				grid[row][col] = 2
				q = append(q, []int{row, col})
				fresh--
			}
		}
		time++
	}
	if fresh == 0 {
		return time
	}
	return -1
}
