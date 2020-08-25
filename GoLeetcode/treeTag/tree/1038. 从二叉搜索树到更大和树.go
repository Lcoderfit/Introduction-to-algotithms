package tree


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var lastVal = 0

func BstToGst(root *TreeNode) *TreeNode {
	if root == nil {
		return  nil
	}
	//这里每次调用需要置为0
	lastVal = 0
	dfs1(root)
	return root
}

func dfs1(node *TreeNode) {
	if node == nil {
		return
	}
	dfs1(node.Right)
	node.Val += lastVal
	lastVal = node.Val
	dfs1(node.Left)
}

//添加函数参数代替全局变量
func dfs2(node *TreeNode, sum *int) {
	if node == nil {
		return
	}
	dfs2(node.Right, sum)
	node.Val += *sum
	*sum = node.Val
	dfs2(node.Left, sum)
}
