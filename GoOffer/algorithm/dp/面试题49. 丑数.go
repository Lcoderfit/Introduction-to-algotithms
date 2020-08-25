package dp


func NthUglyNumber(n int) int {
	dp := make([]int, n)
	p2,p3,p5 := 0, 0, 0
	dp[0] = 1
	for i := 1; i < n; i++ {
		dp[i] = min3(dp[p2]*2,dp[p3]*3,dp[p5]*5)
		if dp[i] == dp[p2]*2 {
			p2++
		}
		if dp[i] == dp[p3]*3 {
			p3++
		}
		if dp[i] == dp[p5]*5 {
			p5++
		}
	}
	return dp[n-1]
}

func min3(a, b, c int) int {
	if a > b {
		a, b = b, a
	}
	if a > c {
		a, c = c, a
	}
	return a
}