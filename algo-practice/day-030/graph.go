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
	for _, n := range g.AdjacencyList[v] {
		if _, ok := g.Visited[n]; !ok {
			g.DFS(n)
		}
	}
}
