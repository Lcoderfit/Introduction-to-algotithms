/*
方法1：动态规划
时间复杂度：O()
空间复杂度：O()

case1:

r:

*/
package medium

func WordBreak(s string, wordDict []string) bool {
	hash := make(map[string]bool, len(wordDict))
	for _, w := range wordDict {
		hash[w] = true
	}
	n := len(s)
	// 因为s[0]需要先判定是否在hash字典中,所以需要扩充一个位置，然后初始化第一个位置为true
	dp := make([]bool, n+1)
	dp[0] = true

	// 扩步法，外层循环控制步幅度，内层循环从头开始循环到步幅尾结束
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			dp[i] = dp[j] && hash[s[j:i]]
		}
	}
	return dp[n]
}
