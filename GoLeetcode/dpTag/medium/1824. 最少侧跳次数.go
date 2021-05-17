/*

方法1：动态规划
时间复杂度：O(n)
空间复杂度：O(n)

case1:

r:

*/
package medium

import "math"

func MinSideJumps(obstacles []int) int {
	if obstacles == nil || len(obstacles) == 0 {
		return 0
	}
	// dp[i]表示从起始点到当前点的第i+1号跑道所需要的最少侧跳数
	dp := make([]int, 3)
	dp[0], dp[1], dp[2] = 1, 0, 1
	for i := 1; i < len(obstacles); i++ {
		// 先将有障碍的跑道设置为最大值，即无法达到
		if obstacles[i] != 0 {
			dp[obstacles[i]-1] = math.MaxInt32
		}
		if obstacles[i] != 1 {
			dp[0] = Min(dp[0], Min(dp[1], dp[2])+1)
		}
		if obstacles[i] != 2 {
			dp[1] = Min(dp[1], Min(dp[0], dp[2])+1)
		}
		if obstacles[i] != 3 {
			dp[2] = Min(dp[2], Min(dp[0], dp[1])+1)
		}
	}
	return Min(dp[0], dp[1], dp[2])
}
