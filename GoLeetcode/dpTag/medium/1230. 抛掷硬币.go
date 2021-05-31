/*
方法1：动态规划
时间复杂度：O()
空间复杂度：O()

case1:

r:

*/
package medium

//
func ProbabilityOfHeads(prob []float64, target int) float64 {
	if prob == nil || len(prob) == 0 {
		return 0
	}
	dp := make([]float64, target+1)
	dp[0] = 1
	for i := 0; i < len(prob); i++ {
		// 这里需要注意一点，如果当前遍历的i+1的大小小于target，则最多只有i+1个硬币正面朝上，不可能有target个硬币正面朝上，所以需要用Min(i+1, target)
		// 第二点：当j==0的时候，表示的含义是没有硬币正面朝上的概率，假设当前遍历到i，dp[0]表示i个硬币均为反面的概率
		for j := Min(target, i+1); j >= 0; j-- {
			if j > 0 {
				dp[j] = dp[j]*(1-prob[i]) + dp[j-1]*prob[i]
			} else {
				dp[j] *= 1 - prob[i]
			}
		}
	}
	return dp[target]
}
