package practice

import "fmt"

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
			// boundry and avaliable slot check
			if nextRow >= 0 && nextRow < rows &&
				nextCol >= 0 && nextCol < cols &&
				maze[nextRow][nextCol] == '.' {
				if nextRow == 0 || nextRow == rows-1 || nextCol == 0 || nextCol == cols-1 {
					return currDistance + 1
				}
				maze[nextRow][nextCol] = '+'
				queue = append(queue, []int{nextRow, nextCol, currDistance + 1})
			}
		}
	}
	return -1
}
