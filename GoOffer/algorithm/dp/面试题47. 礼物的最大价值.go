package dp

import "fmt"

func MaxValue(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	rows, cols := len(grid), len(grid[0])
	dp := Create2DSlice(rows, cols)
	copy(dp, grid)
	fmt.Println(dp)
	for i := 1; i < cols; i++ {
		dp[0][i] = dp[0][i] + dp[0][i-1]
	}
	for j := 1; j < rows; j++ {
		dp[j][0] = dp[j][0] + dp[j-1][0]
	}
	fmt.Println(dp)
	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			dp[i][j] = dp[i][j] + max(dp[i][j-1], dp[i-1][j])
		}
	}
	return dp[rows-1][cols-1]
}