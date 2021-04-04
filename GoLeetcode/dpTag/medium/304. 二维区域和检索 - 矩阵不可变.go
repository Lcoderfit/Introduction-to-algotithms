/*
方法1：二维前缀和
时间复杂度：O(mn)
空间复杂度：O(mn)

case1:

r:

*/
package medium

type NumMatrix struct {
	nums [][]int
}

func ConstructorNumMatrix(matrix [][]int) NumMatrix {
	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1] - dp[i-1][j-1] + matrix[i-1][j-1]
		}
	}
	return NumMatrix{nums: dp}
}

func (nm *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return nm.nums[row2+1][col2+1] + nm.nums[row1][col1] - nm.nums[row1][col2+1] - nm.nums[row2+1][col1]

}
