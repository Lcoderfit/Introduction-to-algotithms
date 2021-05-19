/*

方法1：动态规划
时间复杂度：O(n)
空间复杂度：O(n)

方法2：动态规划+空间优化
时间复杂度：O(n)
空间复杂度：O(1)

case1:

r:

*/
package medium

func NumDecodings(s string) int {
	if len(s) == 0 || s[0] == '0' {
		return 0
	}
	n := len(s)
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		if s[i-2:i] > "26" && s[i-1] != '0' {
			dp[i] = dp[i-1]
		} else if s[i-2:i] > "26" && s[i-1] == '0' {
			return 0
		} else if s[i-2] != '0' && s[i-1] != '0' {
			dp[i] = dp[i-1] + dp[i-2]
		} else if s[i-2] != '0' && s[i-1] == '0' {
			dp[i] = dp[i-2]
		} else if s[i-2] == '0' && s[i-1] != '0' {
			dp[i] = dp[i-1]
		} else {
			return 0
		}
	}
	return dp[n]
}

func NumDecodings3(s string) int {
	if len(s) == 0 || s[0] == '0' {
		return 0
	}
	n := len(s)
	pre1 := 1
	pre2 := 1
	cur := 1
	// 两位分类讨论
	for i := 2; i <= n; i++ {
		if s[i-2:i] > "26" && s[i-1] != '0' {
			cur = pre2
		} else if s[i-2:i] > "26" && s[i-1] == '0' {
			return 0
		} else if s[i-2] != '0' && s[i-1] != '0' {
			cur = pre1 + pre2
		} else if s[i-2] != '0' && s[i-1] == '0' {
			cur = pre1
		} else if s[i-2] == '0' && s[i-1] != '0' {
			cur = pre2
		} else {
			return 0
		}
		pre1 = pre2
		pre2 = cur
	}
	return cur
}

func NumDecodings2(s string) int {
	l := len(s)
	pre1 := 1
	pre2 := 0
	cur := 0
	for i := l - 1; i >= 0; i-- {
		if s[i] == '0' {
			cur = 0
		} else if s[i] >= '3' {
			cur = pre1
		} else if s[i] == '2' {
			cur = pre1
			if i+1 < l && s[i+1] <= '6' {
				cur += pre2
			}
		} else if s[i] == '1' {
			cur = pre1 + pre2
		}
		pre2 = pre1
		pre1 = cur
	}
	return cur
}
