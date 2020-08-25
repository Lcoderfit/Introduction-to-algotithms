/*
面试题26.树的子结构.go
时间复杂度：O(n*m) A为n，B为m，每遍历A一个节点就要调用一次dfs，最坏情况下dfs遍历B的所有节点
空间复杂度：O(n) 添加了一个辅助遍历res，
*/
package tree


func IsSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}

	res := dfs(A, B)
	if !res {
		res = IsSubStructure(A.Left, B)
	}
	if !res {
		res = IsSubStructure(A.Right, B)
	}

	//利用 || 的短路特性可写成
	//return dfs(A,B) || isSubStructure(A.Left,B) || isSubStructure(A.Right,B)
	return res
}

func dfs(a, b *TreeNode) bool {
	if b == nil {
		return true
	}
	if a == nil {
		return false
	}

	if a.Val != b.Val {
		return false
	}
	return dfs(a.Right, b.Right) && dfs(a.Left, b.Left)
}
