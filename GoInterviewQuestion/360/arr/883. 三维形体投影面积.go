package arr


func projectionArea(grid [][]int) int {
	left, front, up := 0, 0, 0
	leftMax, frontMax := 0, 0
	rows, cols := len(grid), len(grid[0])
	for i := 0; i < rows; i++ {
		leftMax = 0
		for j := 0; j < cols; j++ {
			if grid[i][j] != 0 {
				up++
			}
			if leftMax < grid[i][j] {
				leftMax = grid[i][j]
			}
		}
		left += leftMax
	}

	for j := 0; j < cols; j++ {
		frontMax = 0
		for i := 0; i < rows; i++ {
			if frontMax < grid[i][j] {
				frontMax = grid[i][j]
			}
		}
		front += frontMax
	}
	return up+left+front
}