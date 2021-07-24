/*

方法1：
时间复杂度：O(mn)
空间复杂度：O(mn)

case1:

r:


线性dp

LCS：最长公共子序列 longest common subsequence
LIS: 最长上升子序列 longest increasing subsequence

*/
package liner

import (
	. "GoLeetcode/common"
)

func LongestCommonSubsequence(s1 string, s2 string) int {
	if len(s1) == 0 || len(s2) == 0 {
		return 0
	}

	m, n := len(s1), len(s2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = Max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[m][n]
}
