package dp

//dp
func UniquePathsWithObstacles(obstacleGrid [][]int) int {
	rows, cols := len(obstacleGrid), len(obstacleGrid[0])
	if rows == 0 || cols == 0 {
		return 0
	}
	dp := create2DArray(rows, cols)
	for j := 0; j < cols; j++ {
		if obstacleGrid[0][j] == 0 {
			dp[0][j] = 1
		} else {
			break
		}
	}
	for i := 0; i < rows; i++ {
		if obstacleGrid[i][0] == 0 {
			dp[i][0] = 1
		} else {
			break
		}
	}
	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			if obstacleGrid[i][j] == 0 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[rows-1][cols-1]
}

//二：贪心，时间升一维，空间降一维
func UniquePathsWithObstacles2(obstacleGrid [][]int) int {
	rows, cols := len(obstacleGrid), len(obstacleGrid[0])
	if rows == 0 || cols == 0 {
		return 0
	}
	dp := make([]int, cols)
	//初始化为1条路径
	dp[0] = 1
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[j] = 0
				//j>0防止dp[j-1]越界
			} else if j > 0 {
				//dp[j] = dp[j] + dp[j-1],等式右边的dp[j]为上一行的dp指，相当于dp[i-1][j]
				//dp[j-1]相当于dp[i][j-1]
				dp[j] += dp[j-1]
			}
		}
	}
	return dp[cols-1]
}
