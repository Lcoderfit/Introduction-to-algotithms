/*
10-2.青蛙跳台阶
时间复杂度：O(n)
空间复杂度：O(n)
*/
package algorithm

func JumpFloor(number int64) (sum int) {
	if number <= 0 {
		return 0
	}
	if number == 1 {
		return 1
	}
	dp := make([]int, number+1)
	dp[0] = 1
	dp[1] = 1
	for i := int64(2); i < number+1; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[number]
}
