/*

方法1：
时间复杂度：O()
空间复杂度：O()

case1:

r:


有 N 件物品和一个容量是 V 的背包。每件物品只能使用一次。

第 i 件物品的体积是 vi，价值是 wi。

求解将哪些物品装入背包，可使这些物品的总体积不超过背包容量，且总价值最大。

输出 最优选法的方案数。注意答案可能很大，请输出答案模 109+7 的结果。

输入格式
第一行两个整数，N，V，用空格隔开，分别表示物品数量和背包容积。

接下来有 N 行，每行两个整数 vi,wi，用空格隔开，分别表示第 i 件物品的体积和价值。

输出格式
输出一个整数，表示 方案数 模 109+7 的结果。

数据范围
0<N,V≤1000
0<vi,wi≤1000
输入样例
4 5
1 2
2 4
3 4
4 6

两种方案，一种方案是选
1 2
4 6

另一种方案：
2 4
3 4
*/
package knapsack

import "math"

var (
	mod = int(math.Pow(10.0, 9) + 7)
)

func SchemeCountForKnapsack(v int, vi, wi []int) int {
	// dp[j]表示存放物品体积不超过j时所能装物品的最大价值
	dp := make([]int, v+1)
	// result[j]表示物品体积不超过j时所能装物品的价值最大的方案数
	result := make([]int, v+1)
	// 初始化时，一个都不选也算是一种方案
	for i := range result {
		result[i] = 1
	}
	n := len(vi)
	for i := 0; i < n; i++ {
		for j := v; j >= vi[i]; j-- {
			t := dp[j-vi[i]] + wi[i]
			if t > dp[j] {
				dp[j] = t
				result[j] = result[j-vi[i]]
			} else if t == dp[j] {
				result[j] += result[j-vi[i]]
			}
			if result[j] > mod {
				result[j] -= mod
			}
		}
	}
	return result[v]
}
