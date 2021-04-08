/*
方法1：动态规划(二维背包问题)
时间复杂度：O(length*mn)
空间复杂度：O(mn)

case1:
["10","0001","111001","1","0"]
r:
4
*/
package medium

//func FindMaxForm(strs []string, m int, n int) int {
//	length := len(strs)
//	dp := make([][][]int, length+1)
//	for i := range dp {
//		dp[i] = make([][]int, m+1)
//		for j := range dp[i] {
//			dp[i][j] = make([]int, n+1)
//		}
//	}
//
//	for i := 1; i <= length; i++ {
//		one, zero := CountOneAndZero(strs[i-1])
//		for j := 0; j <= m; j++ {
//			for k := 0; k <= n; k++ {
//				if j >= zero && k >= one {
//					// 遍历到第i个元素时，如果不取第i个元素，则dp[i][j][k] = dp[i-1][j][k]
//					// 如果取第i个元素，则dp[i][j][k] = dp[i-1][j-zero][k-one] + 1
//					dp[i][j][k] = Max(dp[i-1][j][k], dp[i-1][j-zero][k-one]+1)
//				} else {
//					dp[i][j][k] = dp[i-1][j][k]
//				}
//			}
//		}
//	}
//
//	for i := range dp {
//		for j := range dp[i] {
//			fmt.Println(dp[i][j])
//		}
//		fmt.Println()
//	}
//
//	return dp[length][m][n]
//}

func FindMaxForm(strs []string, m int, n int) int {
	// 首先，二维背包问题，不需要定义三维的dp（第一维表示位置，第二维表示一个维度的背包，第三维表示另一个维度的背包）
	// 可以只定义一个二维的dp，分别用于代表一个维度的背包;
	// 然后使用三循环，第一层循环遍历每一个元素，第二和第三个循环倒序遍历，分别用于二个维度背包的动态规划
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for _, str := range strs {
		one, zero := CountOneAndZero(str)
		// 对背包的值维度(非位置)进行逆序dp,降低维度
		for i := m; i >= zero; i-- {
			for j := n; j >= one; j-- {
				dp[i][j] = Max(dp[i][j], dp[i-zero][j-one]+1)
			}
		}
	}
	return dp[m][n]
}

func CountOneAndZero(str string) (one int, zero int) {
	for _, c := range str {
		if c == '1' {
			one++
		} else {
			zero++
		}
	}
	return
}
