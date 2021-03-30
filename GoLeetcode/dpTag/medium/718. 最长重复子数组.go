/*
方法1：动态规划
时间复杂度：O(mn)
空间复杂度：O(mn)

case1:
1 2 3 2 1
3 2 1 4 7
r:
3 2 1
*/
package medium

// FindLength 返回数组A和B的最长重复子数组的长度，注意：子数组一定是连续的
func FindLength(A []int, B []int) int {
	m, n := len(A), len(B)
	dp := make([][]int, m+1)
	dp[0] = make([]int, n+1)
	ans := 0
	for i := 1; i <= m; i++ {
		dp[i] = make([]int, n+1)
		for j := 1; j <= n; j++ {
			if A[i-1] == B[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				ans = Max(ans, dp[i][j])
			} else {
				dp[i][j] = 0
			}
		}
	}
	return ans
}
