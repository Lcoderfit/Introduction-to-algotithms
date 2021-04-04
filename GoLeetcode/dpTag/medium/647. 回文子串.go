/*
方法1：动态规划
时间复杂度：O()
空间复杂度：O()

case1:

r:

*/
package medium

import "fmt"

func CountSubstrings(s string) int {
	n := len(s)
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}

	ans := 0
	for i := n-1; i >= 0; i-- {
		dp[i][i] = true
		ans++
		for j := i + 1; j < n; j++ {
			if j - i > 1 {
				dp[i][j] = dp[i+1][j-1] && s[i] == s[j]
			} else {
				dp[i][j] = s[i] == s[j]
			}
			if dp[i][j] {
				ans++
			}
		}
	}
	for _, array := range dp {
		fmt.Println(array)
	}
	return ans
}
