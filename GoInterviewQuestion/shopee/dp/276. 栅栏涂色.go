package dp


//dp
func NumWays(n int, k int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return k
	}
	dp := make([]int, n+1)
	dp[1] = k
	dp[2] = k*k
	for i := 3; i <= n; i++ {
		dp[i] = (k-1)*(dp[i-1] + dp[i-2])
	}
	return dp[n]
}

//贪心
func NumWays1(n int, k int) int {
	if n == 0 || k == 0 {
		return 0
	}
	if n == 1 {
		return k
	}
	if n == 2 {
		return k*k
	}
	ret := 0
	first, second := k, k*k
	for i := 3; i <= n; i++ {
		ret = (first + second)*(k-1)
		first, second = second, ret
	}
	return ret
}