package tree


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ret := 0
	helper(root, &ret)
	return ret
}

func helper(node *TreeNode, ret *int) int {
	if node == nil {
		return 0
	}
	left := helper(node.Left, ret)
	right := helper(node.Right, ret)
	*ret = max(*ret, left + right)
	return 1 + max(left, right)
}