/*
面试题10- II. 青蛙跳台阶问题.go
时间复杂度：O(n)
空间复杂度：O(n)
*/
package dp


func numWays(n int) int {
	if n == 0 {
		return 1
	}
	dp := make([]int, n + 1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		//dp[i] = (dp[i - 1] + dp[i - 2]) % (1e9+7) = (dp[i - 1]%(1e9+7) + dp[i - 2]%(1e9+7)) % (1e9+7)
		dp[i] = (dp[i - 1] + dp[i - 2]) % (1e9+7)
	}

	return dp[n]
}