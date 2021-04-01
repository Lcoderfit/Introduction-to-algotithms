/*
方法1：动态规划
时间复杂度：O(n)
空间复杂度：O(1)

方法2：公式法
时间复杂度：O(n)
空间复杂度：O(1)

case1:

r:

*/
package medium

// NumberOfArithmeticSlices 计算子等差数列的个数
func NumberOfArithmeticSlices(nums []int) int {
	if nums == nil || len(nums) < 3 {
		return 0
	}
	// dp[i]表示当前以元素i结尾的等差数列的个数，假设以nums[i-1]结尾的等差数列有x个，最长的一个为nums[k...i-1]
	// 如果nums[i]-nums[i-1]==nums[i-1]-nums[i-2],则以nums[i]结尾的等差数列最长的一个为nums[k...i], 从k+1到i的个数等于从k到i-1的个数
	// 所以dp[i] = dp[i-1] + 1
	dp := 0
	sum := 0
	for i := 2; i < len(nums); i++ {
		if nums[i]-nums[i-1] == nums[i-1]-nums[i-2] {
			dp = dp + 1
			sum += dp
		} else {
			dp = 0
		}
	}
	return sum
}

func NumberOfArithmeticSlices2(nums []int) int {
	if nums == nil || len(nums) < 3 {
		return 0
	}
	sum := 0
	count := 0
	for i := 2; i < len(nums); i++ {
		// 计算连续的等差数列的长度，然后通过高斯定理计算
		if nums[i]-nums[i-1] == nums[i-1]-nums[i-2] {
			count++
		} else {
			sum += count * (count + 1) / 2
			count = 0
		}
	}
	return sum + count*(count+1)/2
}
