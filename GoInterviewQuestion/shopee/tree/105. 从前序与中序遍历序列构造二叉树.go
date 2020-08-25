package tree


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//树的题目一般都是递归
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	var i int
	//找到中序序号中根节点的位置
	for i = 0; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}
	//控制preorder和inorder分成两部分，分别为左右子树
	root.Left = buildTree(preorder[1:i+1], inorder[:i])
	//inorder需要出去第i个位置的元素，所以切片从i+1开始
	root.Right = buildTree(preorder[i+1:], inorder[i+1:])
	return root
}