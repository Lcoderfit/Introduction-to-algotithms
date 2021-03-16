/*
时间复杂度：O(n^2)
空间复杂度：O(n)
*/
package dp

func Respace(dictionary []string, sentence string) int {
	n := len(sentence)
	//dp[i]表示前i个字符中的最少未识别字符数
	dp := make([]int, n+1)
	//dp[0] = 0
	for i := 1; i <= n; i++ {
		dp[i] = dp[i-1] + 1
		for j := 0; j < i; j++ {
			if find(dictionary, sentence[j:i]) {
				//在string中，标号为j的位置前面有j个字符
				dp[i] = min(dp[i], dp[j])
			}
		}
	}
	return dp[n]
}

func find(dictionary []string, str string) bool {
	for _, s := range dictionary {
		if s == str {
			return true
		}
	}
	return false
}
