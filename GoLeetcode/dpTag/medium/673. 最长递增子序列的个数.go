package medium

func FindNumberOfLIS(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums)+1)
	dp[1] = nums[0]
	maxLen := 1
	ret := 1
	for i := 2; i <= len(nums); i++ {
		for j := 1; j < i; j++ {
			if nums[i-1] > nums[j-1] {
				dp[i] = Max(dp[i], dp[j]+1)
			} else {
				dp[i] = 1
			}

			if maxLen < dp[i] {
				maxLen = dp[i]
				ret = 1
			} else if maxLen == dp[i] && maxLen != 1 {
				ret += 1
			} else if j == i-1 && maxLen == 1 {
				ret += 1
			}
		}
	}
	return ret
}
