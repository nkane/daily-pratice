package leetcode

/*
	1466. Reorder Routes to Make All Paths Lead to the City Zero
*/

/*
Examples:
connections := [][]int{
		{0, 1},
		{1, 3},
		{2, 3},
		{4, 0},
		{4, 5},
	}

adj[{node}] -> [{edge_idx, require_reorder}]
adj[0] -> [{1, 1}, {4, 0}] // 0 -> 1 (requires reordering), 0 -> 4 (does not require reordering)
adj[1] -> [{0, 0}, {3, 1}] // 1 -> 0 (does not require reordering), 1 -> 3 (requires reordering)
adj[2] -> [{3, 1}]         // 2 -> 3 (requires reordering)
adj[3] -> [{1, 0}, {2, 0}] // 3 -> 1 (does not require reordering), 3 -> 2 (does not require reordering)
adj[4] -> [{0, 1}, {5, 1}] // 4 -> 0 (requires reordering), 4 -> 5 (requires reordering)
adj[5] -> [{4, 0}]         // 5 -> 4 (does not require reordering)
*/

func minReorder(n int, connections [][]int) int {
	adj := make([][][2]int, n)
	for _, connection := range connections {
		adj[connection[0]] = append(adj[connection[0]], [2]int{connection[1], 1})
		adj[connection[1]] = append(adj[connection[1]], [2]int{connection[0], 0})
	}
	// Initialize the count of reorders
	count := 0
	// Anonymous function for DFS
	var dfs func(int, int)
	dfs = func(node, parent int) {
		for _, neighbor := range adj[node] {
			if neighbor[0] != parent {
				count += neighbor[1]
				dfs(neighbor[0], node)
			}
		}
	}
	// Start DFS from node 0 with no parent (use -1 as a sentinel value for parent)
	dfs(0, -1)
	return count
}
