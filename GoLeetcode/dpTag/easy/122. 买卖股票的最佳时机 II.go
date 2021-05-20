/*

方法1：贪心算法
时间复杂度：O(n)
空间复杂度：O(1)

case1:

r:

*/
package easy

// 将股票价格看作是折线图，则最大收益为所有的波峰减去该波峰前一个波谷的值的和
func MaxProfitII1(prices []int) int {
	res := 0
	for i := 1; i < len(prices); i++ {
		if prices[i]-prices[i-1] > 0 {
			res += prices[i] - prices[i-1]
		}
	}
	return res
}

// 动态规划，先捋清楚有哪些状态，该问题有两个状态，一个是第几天（股票价格跟当前是哪一天有关）；
// 第二个状态是当前手中是否持有股票，持有设置为1，未持有设置为0
// 则设dp[i][1]表示当前手中持有股票的前i天的最大收益，dp[i][0]表示当前手中未持有股票的前i天的最大收益
// 交替动态规划
// dp[i][0] = Max(dp[i-1][0], dp[i-1][1]+prices[i])
// dp[i][1] = Max(dp[i-1][1], dp[i-1][0]-prices[i])
// 当前状态只与前一个状态有关，所以可以进行空间优化，用两个变量分别表示当前持有股票的最大收益 和 当前未持有股票的最大收益
func MaxProfitII2(prices []int) int {
	dp0, dp1 := 0, -prices[0]
	for i := 1; i < len(prices); i++ {
		dp0 = Max(dp0, dp1+prices[i])
		dp1 = Max(dp1, dp0-prices[i])
	}
	return dp0
}
