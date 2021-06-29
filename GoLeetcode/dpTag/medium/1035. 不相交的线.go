/*

方法1：动态规划
时间复杂度：O(mn)
空间复杂度：O(mn)

case1:
r:

case2:
r:

case3:
r:
*/
package medium

func MaxUncrossedLines(nums1 []int, nums2 []int) int {
	m, n := len(nums1), len(nums2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	//for i := 1; i <= m; i++ {
	//	for j := 1; j <= n; j++ {
	//		if nums1[i-1] == nums2[j-1] {
	//			dp[i][j] = dp[i-1][j-1] + 1
	//		} else {
	//			dp[i][j] = Max(dp[i-1][j], dp[i][j-1])
	//		}
	//	}
	//}
	for i, v := range nums1 {
		for j, w := range nums2 {
			if v == w {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = Max(dp[i][j+1], dp[i+1][j])
			}
		}
	}
	return dp[m][n]
}
