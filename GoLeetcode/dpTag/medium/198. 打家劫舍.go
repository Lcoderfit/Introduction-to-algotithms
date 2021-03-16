/*

方法1：动态规划
时间复杂度：O(n)
空间复杂度：O(n)

方法1：动态规划 + 滚动数组
时间复杂度：O(n)
空间复杂度：O(1)

case1:

r:

*/
package medium

func Rob1(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	n := len(nums)
	dp := make([]int, n+1)
	dp[1] = nums[0]
	// 假设取到最后一个(i),则第i-1个不能取，所以dp[i] = dp[i-2] + nums[i - 1]
	// 假设最后一个不取，则dp[i] = dp[i-1]
	for i := 2; i <= n; i++ {
		dp[i] = Max(dp[i-1], dp[i-2]+nums[i-1])
	}
	return dp[n]
}

func Rob2(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	n := len(nums)
	// 如果当前结果只与前两个结果有关，则可借助两个辅助遍历保留前两个结果，然后每次遍历时更新这两个结果（此为滚动数组）
	first, second := nums[0], Max(nums[0], nums[1])
	for i := 2; i < n; i++ {
		first, second = second, Max(second, first+nums[i])
	}
	return second
}
