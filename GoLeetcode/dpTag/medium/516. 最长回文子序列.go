/*
方法1：动态规划
时间复杂度：O(n^2)
空间复杂度：O(n^2)

case1:

r:

*/
package medium

func LongestPalindromeSubseq(s string) int {
	if len(s) == 0 {
		return 0
	}
	n := len(s)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		dp[i][i] = 1
	}
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = Max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return dp[0][n-1]
}
