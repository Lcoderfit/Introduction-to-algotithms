/*
方法1：动态规划
时间复杂度：O(n*2^n)
空间复杂度：O(n^2)

case1:

r:

*/
package medium

// Partition 一个字符串，分割成若干子串，每个子串都必须为回文串
func Partition(s string) [][]string {
	if len(s) < 1 {
		return [][]string{}
	}
	n := len(s)
	dp := make([][]bool, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]bool, n+1)
		dp[i][i] = true
	}
	for i := n - 1; i > 0; i-- {
		for j := i + 1; j <= n; j++ {
			// 当j = i+1时， dp[i+1][j-1]相当于dp[j][j-1], 横索引比纵索引要更大，默认为false，所以需要单独处理
			if j-i > 1 {
				dp[i][j] = dp[i+1][j-1] && s[i-1] == s[j-1]
			} else {
				dp[i][j] = s[i-1] == s[j-1]
			}
		}
	}
	ans := make([][]string, 0)
	splits := make([]string, 0)
	var dfs func(i int)
	dfs = func(i int) {
		if i == n {
			// 由于将splits多次压入ans时，只是将地址存入ans中，所以每次都需要创建一个新的slice
			ans = append(ans, append([]string{}, splits...))
			return
		}
		for j := i; j < n; j++ {
			// 如果s[i...j]是一个回文串，则向前迈出一步，否则回溯将splits恢复原样
			if dp[i+1][j+1] {
				splits = append(splits, s[i:j+1])
				dfs(j + 1)
				// 去除末尾元素（回溯一步）
				splits = splits[:len(splits)-1]
			}
		}
	}
	dfs(0)
	return ans
}
