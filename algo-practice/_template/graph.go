package practice

type Graph[K comparable, T any] struct {
	Visited       map[K]struct{}
	AdjacencyList map[K][]K
}

func (g *Graph[K, T]) AddEdge(v K, w K) {
}

func (g *Graph[K, T]) DFS(v K) {
}

type Cell struct {
	X    int
	Y    int
	Step int
}

func nearestExit(maze [][]byte, entrance []int) int {
	return -1
}
