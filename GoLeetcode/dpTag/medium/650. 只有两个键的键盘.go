/*

方法1：数学法，因式分解求因式和
时间复杂度：O(根号n)
空间复杂度：O(1)

方法2：动态规划
时间复杂度：O(1)
空间复杂度：O(1)

case1:

r:

*/
package medium

import "math"

// 其实就是求因式分解后，的因数和
func MinSteps1(n int) int {
	if n <= 0 {
		return 0
	}
	d := 2
	res := 0
	for n > 1 {
		for n%d == 0 {
			res += d
			n /= d
		}
		d++
	}
	return res
}

// 例如n为24，则其可以分解为：1*24, 2*12, 3*8, 4*6, 2*2*2*3..... 求的就是分解出来的因数的和的最小值
func MinSteps2(n int) int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = math.MaxInt32
	}
	dp[1] = 0
	for i := 2; i <= n; i++ {
		for j := 1; j < i; j++ {
			if i%j == 0 {
				dp[i] = Min(dp[i], dp[j]+i/j)
			}
		}
	}
	return dp[n]
}
