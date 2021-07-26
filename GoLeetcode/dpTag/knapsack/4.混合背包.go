/*

方法1：
时间复杂度：O()
空间复杂度：O()

case1:

r:


有 N 种物品和一个容量是 V 的背包。

物品一共有三类：

第一类物品只能用1次（01背包）；
第二类物品可以用无限次（完全背包）；
第三类物品最多只能用 si 次（多重背包）；
每种体积是 vi，价值是 wi。

求解将哪些物品装入背包，可使物品体积总和不超过背包容量，且价值总和最大。
输出最大价值。

输入格式
第一行两个整数，N，V，用空格隔开，分别表示物品种数和背包容积。

接下来有 N 行，每行三个整数 vi,wi,si，用空格隔开，分别表示第 i 种物品的体积、价值和数量。

si=−1 表示第 i 种物品只能用1次；
si=0 表示第 i 种物品可以用无限次；
si>0 表示第 i 种物品可以使用 si 次；
*/
package knapsack

type Good1 struct {
	K bool // true表示为01背包，否则为完全背包
	V int
	W int
}

func BlendKnapsack(m int, v, w, s []int) int {
	dp := make([]int, m+1)
	goods := make([]Good1, 0)
	for i := range s {
		if s[i] < 0 {
			goods = append(goods, Good1{K: true, V: v[i], W: w[i]})
		} else if s[i] == 0 {
			goods = append(goods, Good1{K: false, V: v[i], W: w[i]})
		} else {
			// 将多重背包转换为01背包
			t := s[i]
			for k := 1; k <= t; k *= 2 {
				t -= k
				goods = append(goods, Good1{K: true, V: k * v[i], W: k * w[i]})
			}
			if t > 0 {
				goods = append(goods, Good1{K: true, V: t * v[i], W: t * w[i]})
			}
		}
	}

	for _, good := range goods {
		if good.K {
			for j := m; j >= good.V; j-- {
				dp[j] = Max(dp[j], dp[j-good.V]+good.W)
			}
		} else {
			for j := good.V; j <= m; j++ {
				dp[j] = Max(dp[j], dp[j-good.V]+good.W)
			}
		}
	}
	return dp[m]
}
