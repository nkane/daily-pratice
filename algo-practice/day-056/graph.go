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
	rows := len(maze)
	cols := len(maze[0])
	dirs := [][]int{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}
	sr := entrance[0]
	sc := entrance[1]
	maze[sr][sc] = '+'
	q := [][]int{
		{sr, sc, 0},
	}
	for len(q) > 0 {
		cr := q[0][0]
		cc := q[0][1]
		cd := q[0][2]
		q = q[1:]
		for _, d := range dirs {
			nr := cr + d[0]
			nc := cc + d[1]
			if nr >= 0 && nr < rows &&
				nc >= 0 && nc < cols &&
				maze[nr][nc] == '.' {
				if nr == 0 || nr == rows-1 ||
					nc == 0 || nc == cols-1 {
					return cd + 1
				}
				maze[nr][nc] = '+'
				q = append(q, []int{nr, nc, cd + 1})
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
	minutesElapsed := -1
	dirs := [][]int{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}
	for len(q) > 0 {
		r := q[0][0]
		c := q[0][1]
		q = q[1:]
		if r == -1 {
			minutesElapsed++
			if len(q) > 0 {
				q = append(q, []int{-1, -1})
			}
		} else {
			for _, d := range dirs {
				nr := r + d[0]
				nc := c + d[1]
				if rows > nr && nr >= 0 &&
					cols > nc && nc >= 0 {
					if grid[nr][nc] == 1 {
						grid[nr][nc] = 2
						freshOranges--
						q = append(q, []int{nr, nc})
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
	rows := len(grid)
	cols := len(grid[0])
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
				nr := d[0] + r
				nc := d[1] + c
				if nr < 0 || nr == len(grid) ||
					nc < 0 || nc == len(grid[0]) ||
					grid[nr][nc] != 1 {
					continue
				}
				grid[nr][nc] = 2
				q = append(q, []int{nr, nc})
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
