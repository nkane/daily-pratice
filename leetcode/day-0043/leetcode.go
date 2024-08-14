package leetcode

/*
	994: Rotting Oranges
	You are given a m x n grid where each cell can have one of three values:

	- 0, represents an empty cell
	- 1, represents a fresh orange
	- 2, represents a rotten organge

	Every minute, any fresh orange that is 4-directionally adjacent to a rotten orange becomes rotten.

	Return the minimum number of minutes that must elapse until no cell has a fresh orange. If this is
	impossible, return -1.
*/

func orangesRotting(grid [][]int) int {
	queue := [][]int{}
	freshOranges := 0
	rows := len(grid)
	cols := len(grid[0])
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == 2 {
				queue = append(queue, []int{r, c})
			} else if grid[r][c] == 1 {
				freshOranges++
			}
		}
	}
	queue = append(queue, []int{-1, -1})
	minutesElapsed := -1
	directions := [][]int{
		{-1, 0}, // left
		{0, 1},  // up
		{1, 0},  // right
		{0, -1}, // down
	}
	for len(queue) > 0 {
		row, col := queue[0][0], queue[0][1]
		queue = queue[1:]
		if row == -1 {
			minutesElapsed++
			if len(queue) > 0 {
				queue = append(queue, []int{-1, -1})
			}
		} else {
			for _, d := range directions {
				neighborRow, neighborCol := row+d[0], col+d[1]
				if rows > neighborRow && neighborRow >= 0 &&
					cols > neighborCol && neighborCol >= 0 {
					if grid[neighborRow][neighborCol] == 1 {
						grid[neighborRow][neighborCol] = 2
						freshOranges--
						queue = append(queue, []int{neighborRow, neighborCol})
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

func orangesRottingInPlace(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	directions := [][]int{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}
	runRottingProcess := func(timestamp int) bool {
		toBeContinued := false
		for row := range rows {
			for col := range cols {
				if grid[row][col] == timestamp {
					for _, d := range directions {
						nRow, nCol := row+d[0], col+d[1]
						if rows > nRow && nRow >= 0 &&
							cols > nCol && nCol >= 0 {
							if grid[nRow][nCol] == 1 {
								grid[nRow][nCol] = timestamp + 1
								toBeContinued = true
							}
						}
					}
				}
			}
		}
		return toBeContinued
	}
	timestamp := 2
	for runRottingProcess(timestamp) {
		timestamp++
	}
	for _, row := range grid {
		for _, cell := range row {
			if cell == 1 {
				return -1
			}
		}
	}
	return timestamp - 2
}

func orangesRottingNeet(grid [][]int) int {
	q := [][]int{}
	time, fresh := 0, 0
	rows, cols := len(grid), len(grid[0])
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == 1 {
				fresh += 1
			}
			if grid[r][c] == 2 {
				q = append(q, []int{r, c})
			}
		}
	}
	directions := [][]int{
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
			for _, d := range directions {
				row, col := d[0]+r, d[1]+c
				if row < 0 || row == len(grid) ||
					col < 0 || col == len(grid[0]) ||
					grid[row][col] != 1 {
					continue
				}
				grid[row][col] = 2
				q = append(q, []int{row, col})
				fresh--
			}
		}
		time += 1
	}
	if fresh == 0 {
		return time
	}
	return -1
}
