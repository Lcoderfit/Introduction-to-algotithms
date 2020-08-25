/*
面试题28.对称二叉树.go
时间复杂度：O(n)，因为我们遍历整个输入树一次，所以总的运行时间为O(n)O(n)，其中nn是树中结点的总数。
空间复杂度：O(n)，递归调用的次数受树的高度限制。在最糟糕情况下，树是线性的，其高度为O(n)。
因此，在最糟糕的情况下，由栈上的递归调用造成的空间复杂度为O(n)。
*/
package tree


//递归法
func IsSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return Dfs(root.Left, root.Right)
}

func Dfs(a, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if a.Val != b.Val {
		return false
	}

	return Dfs(a.Left, b.Right) && Dfs(a.Right, b.Left)
}

//迭代法
//func IsSymmetric1(root *TreeNode) bool {
//
//}