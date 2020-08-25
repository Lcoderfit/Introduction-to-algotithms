/*
面试题7.重建二叉树.go
时间复杂度：O(n^2) 最坏情况下（左子节点相连的单链），遍历inorder切片，n + (n-1) + (n-2) ... 1 = n*(n+1)/2
空间复杂度：O(n)
*/
package tree


// 递归
func BuildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	head := &TreeNode{Val:preorder[0]}
	var i, val int
	for i, val = range inorder {
		if val == preorder[0] {
			break
		}
	}

	head.Left = BuildTree(preorder[1:i+1], inorder[0:i])
	head.Right = BuildTree(preorder[i+1:], inorder[i+1:])

	return head
}

// 迭代