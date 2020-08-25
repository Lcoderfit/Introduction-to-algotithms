/*
M,N分别为矩阵行列大小,K为字符串 word 长度。

时间复杂度 O(3^KMN) ： 最差情况下，需要遍历矩阵中长度为 K字符串的所有方案，
时间复杂度为 O(3^K)；矩阵中共有 MNMN 个起点，时间复杂度为 O(MN) 。
方案数计算： 设字符串长度为 KK ，搜索中每个字符有上、下、左、右四个方向可以选择，舍弃回头（上个字符）的方向，剩下 3种选择，因此方案数的复杂度为 O(3^K)
空间复杂度 O(K) ： 搜索过程中的递归深度不超过 K，因此系统因函数调用累计使用的栈空间占用 O(K) （因为函数返回后，系统调用的栈空间会释放）。最坏情况下 K = MN，递归深度为 MN，此时系统栈使用 O(MN)的额外空间。
 */
package bfsAndDfsTag


func Exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if dfs(board, word, i, j, 0) {
				return true
			}
		}
	}
	return false
}

func dfs(board [][]byte, word string, i int, j int, k int) bool {
	if i >= len(board) || i < 0 || j >= len(board[0]) || j < 0 || board[i][j] != word[k] {
		return false
	}
	if k == len(word)-1 {
		return true
	}
	tmp := board[i][j]
	board[i][j] = '/'
	res := dfs(board, word, i+1, j, k+1) || dfs(board, word, i-1, j, k+1) ||
		dfs(board, word, i, j+1, k+1) || dfs(board, word, i, j-1, k+1)
	board[i][j] = tmp
	return res
}