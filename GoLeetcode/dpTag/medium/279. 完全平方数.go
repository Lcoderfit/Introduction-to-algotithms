/*
方法1：动态规划
时间复杂度：O(n*根号n)
空间复杂度：O(n)


方法1：贪心+数学公式
时间复杂度：O()
空间复杂度：O()

case1:

r:
*/
package medium

import "math"

// NumSquares 枚举动态规划，dp[i] = Min(dp[i], dp[i-k]+1) (k: [1, 4, 9, .... ], k为枚举值， k为根号n的向下取整值)
func NumSquares(n int) int {
	if n <= 0 {
		return 0
	}

	border := int(math.Floor(math.Sqrt(float64(n))))
	squares := make([]int, border)
	for i := 1; i <= border; i++ {
		squares[i-1] = i * i
	}

	dp := make([]int, n+1)
	dp[0] = 0
	for i := 1; i <= n; i++ {
		dp[i] = math.MaxInt32
		for _, s := range squares {
			if i < s {
				break
			}
			dp[i] = Min(dp[i], dp[i-s]+1)
		}
	}
	return dp[n]
}
