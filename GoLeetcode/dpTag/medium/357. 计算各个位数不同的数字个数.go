/*
方法1：动态规划
时间复杂度：O(n)
空间复杂度：O(n)

方法1：递归
时间复杂度：O(n^2)
空间复杂度：O(1)

case1:

r:


*/
package medium

import "math"

func CountNumbersWithUniqueDigits(n int) int {
	if n == 0 {
		return 0
	}
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 0
	sum := dp[1]
	// dp[i]表示i位数重复的数字个数
	// 前i-1位重复的数字总数为：首先dp[i-1]中的数字全为重复数，所以最后一位可以为0-9中任意数字，即10*dp[i-1]
	//
	// 前i-1位不重复，最后一位与前i-1位重复的数字总数：
	// i-1位数总共有 最高位9种*次高位10种（包括0）*次次高位10种。。。。，即9*int(math.Pow(10, float64(i-2)))
	// i-1位数的数字总数 - i-1位数重复的数字总数==i-1位数不重复的数字总数， 即9*int(math.Pow(10, float64(i-2)))-dp[i-1]
	// 最后一位数只需要与前面i-1位数中任何一位相同即为重复数，即 (9*int(math.Pow(10, float64(i-2)))-dp[i-1])*(i-1)
	for i := 2; i <= n; i++ {
		dp[i] = 10*dp[i-1] + (9*int(math.Pow(10, float64(i-2)))-dp[i-1])*(i-1)
		sum += dp[i]
	}
	return int(math.Pow(10, float64(n))) - sum
}

func CountNumbersWithUniqueDigits2(n int) int {
	if n == 0 {
		return 1
	}
	res := 9
	// 假设有n位，n位数不重复的个数为：最高位9种（1～9）*次高位9种（0～9十种，减去与最高为重复的一种，即9种）*次次高位8种。。。。。。
	for i := 1; i < n; i++ {
		res *= 10 - i
	}
	return res + CountNumbersWithUniqueDigits(n-1)
}
