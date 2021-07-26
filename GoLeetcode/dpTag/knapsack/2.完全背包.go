/*

方法1：
时间复杂度：O()
空间复杂度：O()

case1:

r:


有 N种物品和一个容量是 V 的背包，每种物品都有无限件可用。

第 i 种物品的体积是 vi，价值是 wi。

求解将哪些物品装入背包，可使这些物品的总体积不超过背包容量，且总价值最大。
输出最大价值。

输入格式
第一行两个整数，N，V，用空格隔开，分别表示物品数量和背包容积。

接下来有 N 行，每行两个整数 vi, wi，用空格隔开，分别表示第 i 件物品的体积和价值。

输出格式
输出一个整数，表示最大价值。

数据范围
0<N,V≤1000
0<vi,wi≤1000
*/
package knapsack

func CompleteKnapsack(v int, vi []int, wi []int) int {
	n := len(vi)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, v+1)
	}

	// dp[i][j]表示取前i种物品，最大使用容量不超过j时；物品的最大价值
	for i := 1; i <= n; i++ {
		// 这里必须从0开始，因为当j < vi[i-1]时，k只能取0，也有dp[i][j]=Max(dp[i][j], dp[i-1][j])
		for j := 0; j <= v; j++ {
			// 这里k必须为0，因为当k=0时，dp[i][j]=Max(dp[i][j], dp[i-1][j])
			// 表示第i个物品一个都不取，
			for k := 0; k*vi[i-1] <= j; k++ {
				// 第i种物品，我可以取k个
				dp[i][j] = Max(dp[i][j], dp[i-1][j-k*vi[i-1]]+k*wi[i-1])
			}
		}
	}
	return dp[n][v]
}

func CompleteKnapsack1(v int, vi []int, wi []int) int {
	n := len(vi)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, v+1)
	}

	// dp[i][j]表示取前i种物品，最大使用容量不超过j时；物品的最大价值
	for i := 1; i <= n; i++ {
		// 其实这一步算到最后，dp[v]就相当于是取了k个vi[i-1]
		for j := 0; j <= v; j++ {
			// f[i , j ] = max( f[i-1,j] , f[i-1,j-v]+w ,  f[i-1,j-2*v]+2*w , f[i-1,j-3*v]+3*w , .....)
			// f[i , j-v]= max(            f[i-1,j-v]   ,  f[i-1,j-2*v] + w , f[i-1,j-2*v]+2*w , .....)
			// 推导出： f[i][j]=max(f[i-1][j], f[i,j-v]+w ) 当j >= v时成立
			//
			if j >= vi[i-1] {
				dp[i][j] = Max(dp[i][j], dp[i][j-vi[i-1]]+wi[i-1])
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[n][v]
}

func CompleteKnapsack2(v int, vi []int, wi []int) int {
	n := len(vi)
	// dp[i]表示容积不超过i时所能取得物品的最大价值
	dp := make([]int, v+1)
	for i := 1; i <= n; i++ {
		for j := v; j >= vi[i-1]; j-- {
			// f[i , j ] = max( f[i-1,j] , f[i-1,j-v]+w ,  f[i-1,j-2*v]+2*w , f[i-1,j-3*v]+3*w , .....)
			// f[j] = max( f[j] , f[j-v]+w ,  f[j-2*v]+2*w , f[j-3*v]+3*w , .....)
			// 用上一层的值来更新当前层
			// 这里k可以从0开始，也可以从1开始，因为k=0时，dp[j]=Max(dp[j], dp[j])
			// 不过建议还是从0开始，表示第i个物品一个都不取
			for k := 0; k*vi[i-1] <= j; k++ {
				dp[j] = Max(dp[j], dp[j-k*vi[i-1]]+k*wi[i-1])
			}
		}
	}
	return dp[v]
}

func CompleteKnapsack3(v int, vi []int, wi []int) int {
	n := len(vi)
	dp := make([]int, v+1)
	for i := 1; i <= n; i++ {
		for j := vi[i-1]; j <= v; j++ {
			// f[i][j]=max(f[i-1][j], f[i,j-v]+w ); 用当前层的值来更新当前层
			dp[j] = Max(dp[j], dp[j-vi[i-1]]+wi[i-1])
		}
	}
	return dp[v]
}
