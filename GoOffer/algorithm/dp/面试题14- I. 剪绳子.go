package dp

import "math"


//贪心算法
func CuttingRope(n int) int {
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}
	threeCount, twoCount := 0, 0
	if n % 3 == 0 {
		threeCount = n/3
	} else if n%3 == 1{
		threeCount = n/3 - 1
		twoCount = 2
	} else if n % 3 == 2 {
		threeCount = n/3
		twoCount = 1
	}
	ret := math.Pow(3, float64(threeCount))*math.Pow(2, float64(twoCount))
	return int(ret)
}

//dp
func CuttingRope1(n int) int {
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}
	dp := make([]int, n + 1)
	dp[1] = 1
	dp[2] = 2
	dp[3] = 3
	for i := 4; i <= n; i++ {
		for j := 1; j <= i/2; j++ {
			dp[i] = max(dp[i], dp[j]*dp[i-j])
		}
	}
	return dp[n]
}