/*

方法1：
时间复杂度：O(4*n)
空间复杂度：O(n)

case1:

r:

*/
package medium

func WaysToChange(n int) int {
	if n <= 0 {
		return 0
	}
	costs := []int{1, 5, 10, 25}
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 0; i < 4; i++ {
		for j := 1; j <= n; j++ {
			if j >= costs[i] {
				dp[j] += dp[j-costs[i]]
			}
		}
	}
	return dp[n] % 1000000007
}
