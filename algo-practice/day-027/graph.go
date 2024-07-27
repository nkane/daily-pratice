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
	for i := 0; i < len(g.AdjacencyList[v]); i++ {
		node := g.AdjacencyList[v][i]
		if _, ok := g.Visited[node]; !ok {
			g.DFS(node)
		}
	}
}
