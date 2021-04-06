/*
方法1：动态规划
时间复杂度：O(n)
空间复杂度：O(n)

方法2：动态规划+滚动数组
时间复杂度：O(n)
空间复杂度：O(1)

case1:
3, 6, 5, 1, 8
r:
18
*/
package medium

import (
	"math"
)

func MaxSumDivThree(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	n := len(nums)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, 3)
	}

	// 该初始化非常重要，只有dp[0][0]可以初始化为0，如果dp[0][1]初始化为0则不符合dp[i][j]表示余数为j的最大值
	// 只有当nums[i-1]%3 == 1时,dp[i][j]才会通过dp[i][1] = Max(dp[i-1][1], dp[i-1][0]+nums[i-1])得到大于0的第一个初始化值(否则一直为负数)
	dp[0][0], dp[0][1], dp[0][2] = 0, math.MinInt32, math.MinInt32
	for i := 1; i <= n; i++ {
		if nums[i-1]%3 == 0 {
			dp[i][0] = Max(dp[i-1][0], dp[i-1][0]+nums[i-1])
			dp[i][1] = Max(dp[i-1][1], dp[i-1][1]+nums[i-1])
			dp[i][2] = Max(dp[i-1][2], dp[i-1][2]+nums[i-1])
		} else if nums[i-1]%3 == 1 {
			dp[i][0] = Max(dp[i-1][0], dp[i-1][2]+nums[i-1])
			dp[i][1] = Max(dp[i-1][1], dp[i-1][0]+nums[i-1])
			dp[i][2] = Max(dp[i-1][2], dp[i-1][1]+nums[i-1])
		} else if nums[i-1]%3 == 2 {
			dp[i][0] = Max(dp[i-1][0], dp[i-1][1]+nums[i-1])
			dp[i][1] = Max(dp[i-1][1], dp[i-1][2]+nums[i-1])
			dp[i][2] = Max(dp[i-1][2], dp[i-1][0]+nums[i-1])
		}
	}

	return dp[n][0]
}

func MaxSumDivThree2(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	n := len(nums)
	// dp[i]表示余数为i的最大和
	dp := make([]int, 3)
	for i := 0; i < n; i++ {
		a := dp[0] + nums[i]
		b := dp[1] + nums[i]
		c := dp[2] + nums[i]

		// 假设a%3==x, 则dp[x] = Max(dp[0]+nums[i], a)
		dp[a%3] = Max(dp[a%3], a)
		dp[b%3] = Max(dp[b%3], b)
		dp[c%3] = Max(dp[c%3], c)
	}
	return dp[0]
}
