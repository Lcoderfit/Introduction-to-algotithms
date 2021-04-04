/*

方法1：动态规划
时间复杂度：O(mn)
空间复杂度：O(mn)

方法2：动态规划 + 滚动数组
时间复杂度：O(mn)
空间复杂度：O(n)

q：如何求最长的公共子序列？？

case1:
abcde ace
r:
3

case2:
abc abc
r:
3

case3:
abc def
r:
0
*/
package medium

func LongestCommonSubsequence(text1 string, text2 string) int {
	if len(text1) == 0 || len(text2) == 0 {
		return 0
	}
	m, n := len(text1), len(text2)
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	// 两个字符串之间的动态规划，一般是二维动规，先设置dp[i][j]的含义，
	// 然后找dp[i][j]与之前的dp（例如dp[i-1][j-1],dp[i-1][j],dp[i][j-1]）的关系
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = Max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[m][n]
}