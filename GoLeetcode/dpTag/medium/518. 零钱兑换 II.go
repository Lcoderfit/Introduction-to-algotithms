package medium

func change(amount int, coins []int) int {
	if coins == nil || len(coins) == 0 {
		return 0
	}

	// dp[i]表示coins能构成i的种类数
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, v := range coins {
		for j := 1; j <= amount; j++ {
			if j >= v {
				dp[j] += dp[j-v]
			}
		}
	}
	return dp[amount]
}
