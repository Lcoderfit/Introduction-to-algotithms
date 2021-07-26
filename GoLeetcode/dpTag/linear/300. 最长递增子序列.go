/*

方法1：动态规划
时间复杂度：O(n^2)
空间复杂度：O(n)

方法2：贪心+二分


case1:

r:

*/
package liner

import (
	. "GoLeetcode/common"
)

func lengthOfLIS(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	// 这个换习惯得改，一开始就得用一个变量代替len(nums), 不然后面写起来很麻烦
	n := len(nums)
	dp := make([]int, n+1)
	maxVal := 1
	for i := 1; i <= n; i++ {
		dp[i] = 1
		for j := 1; j < i; j++ {
			if nums[i-1] > nums[j-1] {
				dp[i] = Max(dp[i], dp[j]+1)
			}
			// 逐步dp，最长的值应该要遍历所有dp取最大值，而不是返回dp[n]
			maxVal = Max(maxVal, dp[i])
		}
	}
	return maxVal
}
