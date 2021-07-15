/*

方法1：
时间复杂度：O(n^2)
空间复杂度：O(n)

case1:

r:




有 N 件物品和一个容量是 V 的背包。每件物品只能使用一次。
第 i 件物品的体积是 vi，价值是 wi。

求解将哪些物品装入背包，可使这些物品的总体积不超过背包容量，且总价值最大。
输出最大价值。

输入格式
第一行两个整数，N，V，用空格隔开，分别表示物品数量和背包容积。

接下来有 N 行，每行两个整数 vi,wi，用空格隔开，分别表示第 i 件物品的体积和价值。

输出格式
输出一个整数，表示最大价值。

数据范围
0<N,V≤1000
0<vi,wi≤1000
*/
package knapsack

import "fmt"

func Knapsack01(size int, vi []int, wi []int) int {
	if size == 0 || vi == nil || len(vi) == 0 || wi == nil || len(wi) == 0 {
		return 0
	}
	// dp[i]表示当前背包使用容量i所能装下物体的最大价值
	// dp[0]则表示当前背包使用容量0，即不装任何物体，所以价值为0，初始化dp[0]=0
	//
	// 如果是求体积恰好是m的情况下，求最大价值是多少，则需要初始化dp[0]=0， dp[i]=-INF;
	// 因为这样可以确保所有状态都从dp[0]转移过来的，不是从dp[0]转移dp数值都小于0，
	// 假设有两个物品，体积和重量分别为1 2, 2 4, 背包总容量为4,则"恰好"情况下dp状态为：('-'表示小于0)
	// 4 3 2 1 0 这一行表示背包容量
	// - - - 2 0
	// - 6 4 2 0
	//
	// 如果不要求"恰好"，则dp转移状态:
	// 4 3 2 1 0 这一行表示背包容量, 按从左到右顺序看
	// 2 2 2 2 0
	// 6 6 4 2 0
	dp := make([]int, size+1)
	n := len(vi)
	for i := 1; i <= n; i++ {
		for j := size; j >= vi[i-1]; j-- {
			dp[j] = Max(dp[j], dp[j-vi[i-1]]+wi[i-1])
		}
		// 可以在此处打印出每一次循环后的dp状态，直观看出dp是如何转移的
		//fmt.Println(dp)
	}
	return dp[size]
}

func main() {
	var n, size int
	fmt.Scanln(&n, &size)
	vi := make([]int, n)
	wi := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanln(&vi[i], &wi[i])
	}
	res := Knapsack01(size, vi, wi)
	fmt.Println(res)
}
