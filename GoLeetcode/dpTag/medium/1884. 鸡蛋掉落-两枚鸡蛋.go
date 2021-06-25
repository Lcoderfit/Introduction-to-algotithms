/*
方法1：动态规划
时间复杂度：O(n^2)
空间复杂度：O(n^2)

方法2：二维动态
时间复杂度：O(n^2)
空间复杂度：O(n)

方法2：数学法
时间复杂度：O(1)
空间复杂度：O(1)

case1:

r:

*/
package medium

import "math"

func TwoEggDrop(n int) int {
	if n == 0 {
		return 0
	}

	dp := make([][]int, 3)
	for i := range dp {
		dp[i] = make([]int, n+1)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt32
		}
	}
	dp[1][0], dp[2][0] = 0, 0
	// 如果只剩一个鸡蛋，需要确保一定能测出来哪一层刚好能摔碎，则必须从第一层开始一层一层往上试,
	// 所以对于i层楼，最少需要试i次才能试出来（为什么是i不是i-1，因为题目要求高于楼层f(0<=f<=n)的任何楼层都会碎）
	for i := 1; i <= n; i++ {
		dp[1][i] = i
	}

	// dp[i][j]表示有i个鸡蛋，测试j层楼一定能找到f的所需要的最少操作次数
	for j := 1; j <= n; j++ {
		for k := 1; k <= j; k++ {
			dp[2][j] = Min(Max(dp[1][k-1], dp[2][j-k])+1, dp[2][j])
		}
	}
	return dp[2][n]
}

func TwoEggDrop1(n int) int {
	if n == 0 {
		return 0
	}
	dp := make([]int, n+1)
	// 因为需要计算当j属于[1~i]时dp[i]的最小值，所以dp[i]初始值应该设置为最大值，不然结果就总是0了
	for i := 1; i <= n; i++ {
		dp[i] = math.MaxInt32
		for j := 1; j <= i; j++ {
			dp[i] = Min(dp[i], Max(j-1, dp[i-j])+1)
		}
	}
	return dp[n]
}

func TwoEggDrop2(n int) int {
	if n == 0 {
		return 0
	}
	return int(math.Ceil((math.Sqrt(float64(8*n+1)) - 1) / 2))
}
