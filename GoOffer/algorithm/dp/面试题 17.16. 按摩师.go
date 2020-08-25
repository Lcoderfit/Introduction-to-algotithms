package dp


func Massage(nums []int) int {
	//该数组元素的选取特性是相邻的不能一起选
	//设dp[i]为前i个元素的最长预约时间
	//dp[i] = max(dp[i-1], dp[i-2]+nums[i-1])
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums)+1)
	dp[1] = nums[0]
	for i := 2; i <= len(nums); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i-1])
	}
	return dp[len(nums)]
}