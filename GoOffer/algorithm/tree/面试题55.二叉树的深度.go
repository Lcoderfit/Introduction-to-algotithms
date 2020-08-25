/*
面试题55.二叉树的深度.go
时间复杂度：O(lgn): 第一层需要cn，第二层有计算左子树和右子树深度的两个函数，节点为n/2,单个函数时间：cn/2,
两个函数就是：2*cn/2, 有lgn层
空间复杂度：O(lgn): 递归深度lgn，递归需要用辅助栈
*/
package tree


func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}
