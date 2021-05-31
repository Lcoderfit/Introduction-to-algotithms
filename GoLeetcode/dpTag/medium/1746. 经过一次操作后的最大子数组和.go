/*
方法1：动态规划
时间复杂度：O(n)
空间复杂度：O(1)

case1:

r:

*/
package medium

import "math"

func MaxSumAfterOperation(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	ans, replace, notReplace := math.MinInt32, 0, 0
	// 分为已替换部分和未替换部分，然后下一个已替换部分由上一个未替换部分动规而来；下一个未替换部分由上一个已替换部分动态规划而来
	for _, v := range nums {
		replace, notReplace = Max(v*v, notReplace+v*v, replace+v), Max(v, v+notReplace)
		ans = Max(ans, replace)
	}
	return ans
}
