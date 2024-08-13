package leetcode

/*
	1926: Nearest Exit from Entrance in Maze
	- Read editorial
*/

type Cell struct {
	X    int
	Y    int
	Step int
}

func nearestExit(maze [][]byte, entrance []int) int {
	rows, cols := len(maze), len(maze[0])
	dirs := [][]int{
		{1, 0}, {-1, 0}, {0, 1}, {0, -1},
	}

	// Mark the entrance as visited since it's not an exit.
	startRow, startCol := entrance[0], entrance[1]
	maze[startRow][startCol] = '+'

	// Start BFS from the entrance, and use a slice to store all cells to be visited.
	queue := [][]int{{startRow, startCol, 0}}

	for len(queue) > 0 {
		currRow, currCol, currDistance := queue[0][0], queue[0][1], queue[0][2]
		queue = queue[1:]

		// Check the four neighbor cells for the current cell.
		for _, d := range dirs {
			nextRow, nextCol := currRow+d[0], currCol+d[1]

			// If there's an unvisited empty neighbor:
			if nextRow >= 0 && nextRow < rows && nextCol >= 0 && nextCol < cols && maze[nextRow][nextCol] == '.' {
				// If this empty cell is an exit, return distance + 1.
				if nextRow == 0 || nextRow == rows-1 || nextCol == 0 || nextCol == cols-1 {
					return currDistance + 1
				}

				// Otherwise, add this cell to the queue and mark it as visited.
				maze[nextRow][nextCol] = '+'
				queue = append(queue, []int{nextRow, nextCol, currDistance + 1})
			}
		}
	}

	// If we finish iterating without finding an exit, return -1.
	return -1
}
