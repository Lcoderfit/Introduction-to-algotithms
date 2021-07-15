/*

方法1：
时间复杂度：O()
空间复杂度：O()

case1:

r:

*/
package easy

//func MaxProfitIV1(k int, prices []int) int {
//	if prices == nil || len(prices) == 0 || k <= 0 {
//		return 0
//	}
//	sellers := make([]int, k)
//	bullers := make([]int, k)
//	dp := make([][]int, len(prices))
//	for i := 0; i < len(prices); i++ {
//		dp[i] = make([]int, k)
//	}
//	bullers[0] = -prices[0]
//	for i := 1; i < len(prices); i++ {
//		dp[i][0] = Max(dp[i-1][0], dp[i-1][1])
//	}
//	return dp[0][0]
//}
