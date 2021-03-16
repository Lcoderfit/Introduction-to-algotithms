/*

方法1：动态规划
时间复杂度：O(mn)
空间复杂度：O(mn)

case1:
4 5
1 0 1 0 0
1 0 1 1 1
1 1 1 1 1
1 0 0 1 0

r:
4

case2:
0 1
1 0
r:
0
*/
package medium

func MaximalSquare1(matrix [][]byte) int {
	if matrix == nil || len(matrix[0]) == 0 {
		return 0
	}
	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		dp[i][0] = int(matrix[i][0]) - 48
	}
	for j := 0; j < n; j++ {
		dp[0][j] = int(matrix[0][j]) - 48
	}

	ans := 0
	if (matrix[0][0] == '1') || (len(matrix) > 1) &&
		(matrix[1][0] == '1') || (len(matrix[0]) > 1 && matrix[0][1] == '1') {
		ans = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == '1' {
				dp[i][j] = MinArray(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
				if dp[i][j] > ans {
					ans = dp[i][j]
				}
			} else {
				dp[i][j] = 0
			}
		}
	}
	return ans * ans
}

func MaximalSquare2(matrix [][]byte) int {
	if matrix == nil || len(matrix[0]) == 0 {
		return 0
	}
	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
				// 因为dp所有元素默认为0，所以只需要考虑为1的情况
				// 直接循环，然后通过该if语句，可以省略单独对第一行和第一列的处理
				if i == 0 || j == 0 {
					dp[i][j] = 1
				} else {
					dp[i][j] = Min(dp[i-1][j-1], Min(dp[i-1][j], dp[i][j-1])) + 1
				}
			}
			ans = Max(ans, dp[i][j])
		}
	}
	return ans * ans
}

func MinArray(array ...int) int {
	ans := array[0]
	for i := 1; i < len(array); i++ {
		if ans > array[i] {
			ans = array[i]
		}
	}
	return ans
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
