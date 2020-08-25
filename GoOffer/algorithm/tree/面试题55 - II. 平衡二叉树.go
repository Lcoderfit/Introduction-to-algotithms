/*
面试题55 - II. 平衡二叉树
时间复杂度：O(n^2)
空间复杂度：O(nlgn)
*/
package tree


func IsBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if abs(treeDepth(root.Left), treeDepth(root.Right)) > 1 {
		return false
	}
	return IsBalanced(root.Left) && IsBalanced(root.Right)
}

func treeDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + max(treeDepth(root.Left), treeDepth(root.Right))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
