package medium

func combinationSum4(nums []int, target int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	// dp[i]表示能构成i的所有种类数
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 1; i <= target; i++ {
		// 遍历到v时，表示当天取v，有dp[i-v]种方法
		for _, v := range nums {
			if i >= v {
				dp[i] += dp[i-v]
			}
		}
	}
	return dp[target]
}
