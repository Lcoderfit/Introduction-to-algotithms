/*

方法1：
时间复杂度：O()
空间复杂度：O()

case1:

r:


有 N 件物品和一个容量是 V 的背包，背包能承受的最大重量是 M。

每件物品只能用一次。体积是 vi，重量是 mi，价值是 wi。

求解将哪些物品装入背包，可使物品总体积不超过背包容量，总重量不超过背包可承受的最大重量，且价值总和最大。
输出最大价值。

输入格式
第一行两个整数，N，V,M，用空格隔开，分别表示物品件数、背包容积和背包可承受的最大重量。

接下来有 N 行，每行三个整数 vi,mi,wi，用空格隔开，分别表示第 i 件物品的体积、重量和价值。

输出格式
输出一个整数，表示最大价值。
*/
package knapsack

func TwoDimensionalExpenseKnapsack(v, m int, vi, mi, wi []int) int {
	// 定义二维数组，一维为体积，二维表示重量
	dp := make([][]int, v+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}

	n := len(vi)

	for k := 0; k < n; k++ {
		for i := v; i >= vi[k]; i-- {
			for j := m; j >= mi[k]; j-- {
				dp[i][j] = Max(dp[i][j], dp[i-vi[k]][j-mi[k]]+wi[k])
			}
		}
	}
	return dp[v][m]
}
