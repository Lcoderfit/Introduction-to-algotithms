package Random


func LukyRandom(n int) float64 {
	if n > 990 {
		return 1.0
	}
	dp := make([]float64, n + 1)
	for i := 1; i <= n; i++ {
		dp[i] = 1.0 - (1.0-dp[i-1])*(float64(990-i+1)/float64(1000-i+1))
	}
	return dp[n]
}