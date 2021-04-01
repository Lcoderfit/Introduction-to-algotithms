/*
方法1：动态规划
时间复杂度：O(n^2)
空间复杂度：O(n^2)

case1:
	triangle1 := [][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	}
r:
11

方法2：动态规划 + 空间优化
时间复杂度：O(n^2)
空间复杂度：O(n)

*/
package medium

import "fmt"

func MinimumTotal(triangle [][]int) int {
	n := len(triangle)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	dp[0][0] = triangle[0][0]
	for i := 1; i < n; i++ {
		dp[i][0] = dp[i-1][0] + triangle[i][0]
		for j := 1; j < i; j++ {
			dp[i][j] = Min(dp[i-1][j], dp[i-1][j-1]) + triangle[i][j]
		}
		dp[i][i] = dp[i-1][i-1] + triangle[i][i]
	}
	ans := Min(dp[n-1]...)
	return ans
}

// MinimumTotal2
func MinimumTotal2(triangle [][]int) int {
	n := len(triangle)
	dp := make([]int, n)
	dp[0] = triangle[0][0]
	temp := 0
	for i := 1; i < n; i++ {
		temp = dp[0]
		dp[0] += triangle[i][0]
		for j := 1; j < i; j++ {
			temp, dp[j] = dp[j], Min(temp, dp[j])+triangle[i][j]
		}
		dp[i] += temp + triangle[i][i]
		fmt.Println(dp)
	}
	return Min(dp...)
}
