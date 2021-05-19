/*
方法1：动态规划
时间复杂度：O(n*target)
空间复杂度：O(n*target)

方法1：动态规划+空间优化
时间复杂度：O(n*target)
空间复杂度：O(target)

case1:
1 5 11 5
r:
true
*/
package medium

func CanPartition(nums []int) bool {
	if nums == nil || len(nums) <= 1 {
		return false
	}
	// 可以转换为0-1背包问题，背包容量为nums总和的一半
	target, max := 0, 0
	for _, v := range nums {
		target += v
		if v > max {
			max = v
		}
	}
	if target&1 == 1 {
		return false
	}
	target = target / 2
	if max > target {
		return false
	}
	n := len(nums)

	// 根据题目意思对dp的元素值类型进行变化，如果题目求是否满足某条件，则设置为布尔类型
	// dp[i][j]表示从前i个元素中是否能取出某些元素使得他们的和为j
	dp := make([][]bool, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]bool, target+1)
	}

	// 不管取到第几个数字，如果我一个数字都不取，总能满足和为0的情况
	for i := 0; i <= n; i++ {
		dp[i][0] = true
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= target; j++ {
			// i表示当前取到第i个元素
			if j >= nums[i-1] {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i-1]]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[n][target]
}

func CanPartition2(nums []int) bool {
	if nums == nil || len(nums) <= 1 {
		return false
	}
	target, max := 0, 0
	for _, v := range nums {
		target += v
		if v > max {
			max = v
		}
	}

	// 总和不能为奇数，且最大值不能大于一半，否则无法分
	if target%2 == 1 {
		return false
	}
	target /= 2
	if max > target {
		return false
	}

	dp := make([]bool, target+1)
	dp[0] = true
	for i := 0; i < len(nums); i++ {
		// 需要倒序动规，因为如果顺序动态规划，dp[j-nums[i]]会先被更新，保存的就不是上一次循环的结果，所以得到的dp[j]也就不是这一次循环的结果
		// 倒序动规，dp[j-nums[i]]保存的就是上一次循环的结果
		for j := target; j >= nums[i]; j-- {
			dp[j] = dp[j] || dp[j-nums[i]]
		}
	}
	return dp[target]
}
