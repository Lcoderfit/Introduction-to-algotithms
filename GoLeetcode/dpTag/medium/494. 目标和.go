/*
方法1：动态规划
时间复杂度：O()
空间复杂度：O()

case1:
1 1 1 1 1
3
r:
5
*/
package medium

//func Solution(nums []int, s int) int {
//	if nums == nil || len(nums) == 0 {
//		return 0
//	}
//
//	n := len(nums)
//	// 题目说了初始的数组的和不会超过1000,而元素值进行减操作可能会有负值,所以需要对值进行+1000的操作,以保证值>=0
//	dp := make([][]int, n+1)
//	for i := range dp {
//		dp[i] = make([]int, 2001)
//	}
//
//	for i := 1; i <= n; i++ {
//		for j := 1; j <= 2000; j++ {
//			if j >= nums[i-1] && j+nums[i-1] <= 2000 {
//				dp[i][j] = dp[i-1][j-nums[i-1]] + dp[i-1][j+nums[i-1]]
//			} else if j >= nums[i-1] {
//				dp[i][j] = dp[i-1][j-nums[i-1]]
//			} else if j+nums[i-1] <= 2000 {
//				dp[i][j] = dp[i-1][j+nums[i-1]]
//			} else {
//				dp
//			}
//		}
//	}
//}

func findTargetSumWays(nums []int, S int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	// 原问题等同于： 找到nums一个正子集和一个负子集，使得总和等于target
	//
	// 我们假设P是正子集，N是负子集 例如： 假设nums = [1, 2, 3, 4, 5]，target = 3，一个可能的解决方案是+1-2+3-4+5 = 3 这里正子集P = [1, 3, 5]和负子集N = [2, 4]
	//
	// 那么让我们看看如何将其转换为子集求和问题：
	//
	//                  sum(P) - sum(N) = target
	// sum(P) + sum(N) + sum(P) - sum(N) = target + sum(P) + sum(N)
	//                       2 * sum(P) = target + sum(nums)
	// 因此，原来的问题已转化为一个求子集的和问题： 找到nums的一个子集 P，使得sum(P) = (target + sum(nums)) / 2
	// target + sum(nums)必须为偶数
	sum := 0
	for _, v := range nums {
		sum += v
	}
	target := sum + S
	// 所有数总和不能比S还小
	if sum < S || target&1 == 1 {
		return 0
	}
	target = target / 2
	// dp[i]代表的含义是从nums中取数相加和为i时有多少种取法, dp[0]表示取数相加和为0的方法数，也就是一个数也不取，只有这一种方法
	dp := make([]int, target+1)
	dp[0] = 1
	// 转换为01背包问题，求和为target的方法数
	for _, v := range nums {
		for j := target; j >= v; j-- {
			dp[j] += dp[j-v]
		}
	}
	return dp[target]
}
