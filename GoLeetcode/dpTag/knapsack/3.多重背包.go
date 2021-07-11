/*

方法1：
时间复杂度：O()
空间复杂度：O()

case1:

r:


有 N 种物品和一个容量是 V 的背包。

第 i 种物品最多有 si 件，每件体积是 vi，价值是 wi。

求解将哪些物品装入背包，可使物品体积总和不超过背包容量，且价值总和最大。
输出最大价值。

输入格式
第一行两个整数，N，V，用空格隔开，分别表示物品种数和背包容积。

接下来有 N 行，每行三个整数 vi,wi,si，用空格隔开，分别表示第 i 种物品的体积、价值和数量。

输出格式
输出一个整数，表示最大价值。

数据范围
0<N,V≤100
0<vi,wi,si≤100
*/
package knapsack

func MultipleChoiceKnapsack(m int, v, w, s []int) int {
	dp := make([]int, m+1)
	n := len(v)
	for i := 0; i < n; i++ {
		for j := m; j >= v[i]; j-- {
			for k := 1; k <= s[i] && k*v[i] <= j; k++ {
				dp[j] = Max(dp[j], dp[j-k*v[i]]+k*w[i])
			}
		}
	}
	return dp[m]
}

type Good struct {
	V int
	W int
}

func MultipleChoiceKnapsack1(m int, v, w, s []int) int {
	dp := make([]int, m+1)
	n := len(v)
	goods := make([]Good, 0)

	// 每种物品有si个，就根据快速幂将每种的si个物品拆分为log2(si)种物品(向上取整)
	for i := 0; i < n; i++ {
		t := s[i]
		for k := 1; k <= t; k *= 2 {
			t -= k
			goods = append(goods, Good{V: k * v[i], W: k * w[i]})
		}
		if t > 0 {
			goods = append(goods, Good{V: t * v[i], W: t * w[i]})
		}
	}

	for _, good := range goods {
		for j := m; j >= good.V; j-- {
			dp[j] = Max(dp[j], dp[j-good.V]+good.W)
		}
	}
	return dp[m]
}
