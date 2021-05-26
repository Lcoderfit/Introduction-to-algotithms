/*

方法1：
时间复杂度：O()
空间复杂度：O()

case1:

r:

*/
package medium

// 连续的子数组和
func CeckSubarraySum(nums []int, k int) bool {
	if nums == nil || len(nums) == 0 || k <= 0 {
		return false
	}
	dp := make(map[int]int, 0)
	dp[0] = -1
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		sum %= k

		if _, ok := dp[sum]; ok {
			// 这个判断必须放在里面，而不能放在ok后面用&&连接，因为就算不符合这个条件，至少sum这个key是存在的，不需要再更新sum到i的映射
			// 而如果放到ok后面，则即使sum这个key已经存在了，但是不满足i-dp[sum]>1的话，也会更新sum:i的映射
			if i-dp[sum] > 1 {
				return true
			}
		} else {
			dp[sum] = i
		}
	}
	return false
}
