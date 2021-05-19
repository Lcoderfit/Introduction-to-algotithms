/*
方法1：动态规划
时间复杂度：O(n)
空间复杂度：O(n)

case1:

r:

*/
package medium

// 集合动态规划
func LongestSubSequence(arr []int, difference int) int {
	if arr == nil || len(arr) == 0 {
		return 0
	}
	dp := make(map[int]int)
	res := 1
	for _, v := range arr {
		dp[v] = Max(dp[v], dp[v-difference]+1)
		res = Max(res, dp[v])
	}
	return res
}

// 该方法超时，不予采用(时间复杂度为O(n^2))
func LongestSubSequence2(arr []int, difference int) int {
	if arr == nil || len(arr) == 0 {
		return 0
	}
	n := len(arr)
	// dp[i]表示以arr[i]结尾的最长定差子序列
	dp := make([]int, n)
	dp[0] = 1
	res := 1
	for i := 1; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if arr[i]-arr[j] == difference {
				dp[i] = Max(dp[i], dp[j]+1)
			}
		}
		res = Max(res, dp[i])
	}
	return res
}
