package leetcode

/*
	399. Evaluate Division
	- neetcode explaination:
	  https://www.youtube.com/watch?v=Uei1fwDoyKk
*/

type Pair struct {
	Variable string
	Value    float64
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	adj := map[string][]Pair{}
	for i, eq := range equations {
		a := eq[0]
		b := eq[1]
		if _, exist := adj[a]; !exist {
			adj[a] = make([]Pair, 0)
		}
		adj[a] = append(adj[a], Pair{
			Variable: b,
			Value:    values[i],
		})
		if _, exist := adj[b]; !exist {
			adj[b] = make([]Pair, 0)
		}
		adj[b] = append(adj[b], Pair{
			Variable: a,
			Value:    1.0 / values[i],
		})
	}
	result := []float64{}
	for _, q := range queries {
		v := bfs(q[0], q[1], adj)
		result = append(result, v)
	}
	return result
}

func bfs(src string, target string, adj map[string][]Pair) float64 {
	if _, ok := adj[src]; !ok {
		return -1
	}
	if _, ok := adj[target]; !ok {
		return -1
	}
	visited := map[string]struct{}{}
	q := []Pair{
		{
			Variable: src,
			Value:    1,
		},
	}
	visited[src] = struct{}{}
	for len(q) > 0 {
		n := q[0]
		q = q[1:]
		if n.Variable == target {
			return n.Value
		}
		for _, neigh := range adj[n.Variable] {
			if _, seen := visited[neigh.Variable]; !seen {
				q = append(q, Pair{
					Variable: neigh.Variable,
					Value:    n.Value * neigh.Value,
				})
				visited[neigh.Variable] = struct{}{}
			}
		}
	}
	return -1
}
