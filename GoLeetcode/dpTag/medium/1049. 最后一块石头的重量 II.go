package medium

func lastStoneWeightII(stones []int) int {
	if stones == nil || len(stones) == 0 {
		return 0
	}
	sum := 0
	for _, v := range stones {
		sum += v
	}
	target := sum / 2
	m := len(stones)

	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, target+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= target; j++ {
			dp[i][j] = dp[i-1][j]
			if j >= stones[i-1] {
				dp[i][j] = Max(dp[i][j], dp[i-1][j-stones[i-1]]+stones[i-1])
			}
		}
	}
	return sum - 2*dp[m][target]
}

func lastStoneWeightII1(stones []int) int {
	if stones == nil || len(stones) == 0 {
		return 0
	}
	sum := 0
	for _, v := range stones {
		sum += v
	}
	target := sum / 2
	m := len(stones)

	// 因为总和最大只有target，所以dp最大索引为target即可
	dp := make([]int, target+1)
	for i := 1; i <= m; i++ {
		// 1.计算当前层时，需要使用上一层的数据，所以需要逆序遍历
		// 2.如果j < stones[i-1],则dp[i][j] = dp[i-1][j],即当前层dp[j]等于上一层的dp[j](令dp[j]的值不变即可), 所以for循环条件设为j >= stones[i-1]
		// 3.如果j >= stones[i-1], dp[i][j] = Max(dp[i][j], dp[i-1][j-stones[i-1]]+stones[i-1]) (当前元素取与不取的问题);
		// 即当前层dp[j]取上一层dp[j]与dp[j-stones[i-1]]+stones[i-1]的最大值
		for j := target; j >= stones[i-1]; j-- {
			dp[j] = Max(dp[j], dp[j-stones[i-1]]+stones[i-1])
		}
	}
	return sum - 2*dp[target]
}
