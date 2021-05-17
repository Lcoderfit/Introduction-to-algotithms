/*
方法1：动态规划
时间复杂度：O(n)
空间复杂度：O(1)

case1:

r:

*/
package medium

import "math"

func maxSumAfterOperation(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	ans, replace, notReplace := math.MinInt32, 0, 0
	for _, v := range nums {
		replace, notReplace = Max(v*v, notReplace+v*v, replace+v), Max(v, v+replace)
		ans = Max(ans, replace)
	}
	return ans
}
