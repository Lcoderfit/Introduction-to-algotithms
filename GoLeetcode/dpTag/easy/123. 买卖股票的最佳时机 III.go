/*

方法1：
时间复杂度：O()
空间复杂度：O()

case1:

r:

本题难度为hard，但是由于是买卖股票系列，所以放在一起
*/
package easy

// leetcode刷题插件
func MaxProfitIII1(prices []int) int {
	if prices == nil || len(prices) <= 1 {
		return 0
	}
	// 由于在同一天买入卖出收益为0，对结果无影响，所以buy2初始化为-prices[0](表示在同一天买入卖出再买入)
	seller1, seller2 := 0, 0
	buy1, buy2 := -prices[0], -prices[0]
	for i := 1; i < len(prices); i++ {
		// 注意，由于buy1为第一次买入，在次之前都没有买卖动作，所以buy1为-prices[i]
		buy1, seller1 = Max(buy1, -prices[i]), Max(seller1, buy1+prices[i])
		buy2, seller2 = Max(buy2, seller1-prices[i]), Max(seller2, buy2+prices[i])
	}
	return seller2
}
