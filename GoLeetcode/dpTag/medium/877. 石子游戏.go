/*
方法1：分析法
时间复杂度：O(1)
空间复杂度：O(1)

case1:

r:

*/
package medium

func StoneGame(piles []int) bool {
	// 因为堆数为偶数堆，所以先手的人总能满足取到偶数序号或者奇数序号的堆，而又因为石头总数为奇数，所以偶数序号堆与奇数序号堆的石头总数一定不相等
	// 且肯定有一个大，一个小，所以先手的人只需要计算偶数序号堆和奇数序号堆中哪一种的石头总数更大，然后取相应的堆就一定能保证赢
	return true
}

// StoneGame2 石头游戏动态规划解法
func StoneGame2(piles []int) bool {
	n := len(piles)
	// dp[i][j]表示对石头堆piles[i...j]进行操作,先手的人最终得到的分数与后手的人最终得到的分数差(两个人都发挥最佳水平)
	// dp[0][n-1]即为对piles[0....n-1]进行操作，先手与后手的分数差
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		// 如果只对一堆石头进行操作，则先手最终获得的分数就是这堆石头的个数
		dp[i][i] = piles[i]
	}

	// 如果i>j，则这堆石头没有意义，因为用到了dp[i+1][j],所以外层循环需要倒序遍历
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			// 如果先手为A,则他可以取piles[i]或者piles[j],取完后剩下的石头,先手一定是B(因为是轮流取,A取完之后就是B取)
			// 所以对piles[i...j]进行操作,先手如果取piles[i], 则dp[i+1][j]为B作为先手对piles[i+1...j]操作时,两人的分数差(B-A)
			// 如果先手取piles[j], 则dp[i][j-1]为B作为先手对piles[i...j-1]操作时,两人的分数差
			dp[i][j] = Max(piles[i]-dp[i+1][j], piles[j]-dp[i][j-1])
		}
	}
	// 亚历山大分数大于0则表示亚历山大获胜
	return dp[0][n-1] > 0
}
