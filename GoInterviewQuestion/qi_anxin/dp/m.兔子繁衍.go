package dp


func CountRabbitNum(n int) int {
	if n < 5 {
		return 1
	}
	dp := make([]int, n+1)
	for i := 0; i < 5; i++ {
		dp[i] = 1
	}
	for i := 5; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-4]
	}
	return dp[n]
}