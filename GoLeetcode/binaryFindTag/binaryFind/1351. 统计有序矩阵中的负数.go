package binaryFind
/*
1351. 统计有序矩阵中的负数.py
方法1： 二分查找
时间复杂度：O(mlogn)
空间复杂度：O(1)

法2：倒序遍历
时间复杂度：O(m)
空间复杂度：O(1)

case:
4 4
4 3 2 -1
3 2 1 -1
1 1 -1 -2
-1 -1 -2 -3
 */

func CountNegatives(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	pos := 0
	res := 0
	for i := m - 1; i >= 0; i-- {
		for j := pos; j < n; j++ {
			if grid[i][j] < 0 {
				pos = j
				res += n - pos
				break
			}
		}
	}
	return res
}
