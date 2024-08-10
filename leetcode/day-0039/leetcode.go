package leetcode

import "container/list"

/*
	547. Number of Provinces
	The are n cities, some of them are connected, while some are not. If a city a is connected
	directly with city b, and city b is connected directly with city c, then city a is connected
	indirectly with city c.

	A province is a group of directly or indirectly connected cities and no other cities outside
	of the group.

	You are given an n x n matrix isConnected where isConnected[i][j] = 1 if the ith city and the
	jth city are directly connected, and isConnected[i][j] = 0 otherwise.
*/

func findCircleNum(isConnected [][]int) int {
	adj := map[int][]int{}
	for i := 0; i < len(isConnected); i++ {
		for j := 0; j < len(isConnected); j++ {
			val := isConnected[i][j]
			if val == 1 {
				if _, ok := adj[i]; !ok {
					adj[i] = []int{
						j,
					}
				} else {
					adj[i] = append(adj[i], j)
				}
			}
		}
	}
	visited := map[int]struct{}{}
	count := 0
	for k := range adj {
		count += dfs(k, visited, adj)
	}
	return count
}

func dfs(currentNode int, visited map[int]struct{}, adj map[int][]int) int {
	if _, ok := visited[currentNode]; ok {
		return 0
	}
	visited[currentNode] = struct{}{}
	neigbors := adj[currentNode]
	for _, n := range neigbors {
		dfs(n, visited, adj)
	}
	return 1
}

// findCircleNum returns the number of connected components in the graph.
func findCircleNum_alt_dfs(isConnected [][]int) int {
	n := len(isConnected)
	visited := make([]bool, n)

	// dfs performs a Depth-First Search starting from the node.
	var dfs func(int)
	dfs = func(node int) {
		visited[node] = true
		for i := 0; i < n; i++ {
			if isConnected[node][i] == 1 && !visited[i] {
				dfs(i)
			}
		}
	}

	numberOfComponents := 0
	for i := 0; i < n; i++ {
		if !visited[i] {
			numberOfComponents++
			dfs(i)
		}
	}
	return numberOfComponents
}

// Solution struct for containing methods related to the problem
type Solution struct{}

// bfs performs a Breadth-First Search starting from the given node.
func (s *Solution) bfs(node int, isConnected [][]int, visited []bool) {
	queue := list.New()
	queue.PushBack(node)
	visited[node] = true

	for queue.Len() > 0 {
		currentNode := queue.Remove(queue.Front()).(int)

		for i := 0; i < len(isConnected); i++ {
			if isConnected[currentNode][i] == 1 && !visited[i] {
				queue.PushBack(i)
				visited[i] = true
			}
		}
	}
}

// findCircleNum returns the number of connected components in the graph.
func (s *Solution) findCircleNum_alt_bfs(isConnected [][]int) int {
	n := len(isConnected)
	visited := make([]bool, n)
	numberOfComponents := 0

	for i := 0; i < n; i++ {
		if !visited[i] {
			numberOfComponents++
			s.bfs(i, isConnected, visited)
		}
	}

	return numberOfComponents
}
