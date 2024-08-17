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
	sR, sC := entrance[0], entrance[1]
	maze[sR][sC] = '+'
	q := [][]int{
		{sR, sC, 0},
	}
	for len(q) > 0 {
		cR, cC, cD := q[0][0], q[0][1], q[0][2]
		q = q[1:]
		for _, d := range dirs {
			nR, nC := cR+d[0], cC+d[1]
			// boundary and availabilty check
			if nR >= 0 && nR < rows &&
				nC >= 0 && nC < cols &&
				maze[nR][nC] == '.' {
				// reach an exit
				if nR == 0 || nR == rows-1 ||
					nC == 0 || nC == cols-1 {
					return cD + 1
				}
				maze[nR][nC] = '+'
				q = append(q, []int{nR, nC, cD + 1})
			}
		}
	}
	return -1
}

func orangesRotting(grid [][]int) int {
	q := [][]int{}
	freshOranges := 0
	rows := len(grid)
	cols := len(grid[0])
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == 2 {
				q = append(q, []int{r, c})
			} else if grid[r][c] == 1 {
				freshOranges++
			}
		}
	}
	q = append(q, []int{-1, -1})
	minutesElasped := -1
	dirs := [][]int{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}
	for len(q) > 0 {
		r, c := q[0][0], q[0][1]
		q = q[1:]
		if r == -1 {
			minutesElasped++
			if len(q) > 0 {
				q = append(q, []int{-1, -1})
			}
		} else {
			for _, d := range dirs {
				nR, nC := r+d[0], c+d[1]
				if rows > nR && nR >= 0 &&
					cols > nC && nC >= 0 {
					if grid[nR][nC] == 1 {
						grid[nR][nC] = 2
						freshOranges--
						q = append(q, []int{nR, nC})
					}
				}
			}
		}
	}
	if freshOranges == 0 {
		return minutesElasped
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
	dirs := [][]int{
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
			for _, d := range dirs {
				r, c := d[0]+r, d[1]+c
				if r < 0 || r == len(grid) ||
					c < 0 || c == len(grid[0]) ||
					grid[r][c] != 1 {
					continue
				}
				grid[r][c] = 2
				q = append(q, []int{r, c})
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
