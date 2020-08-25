package tree


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func ConvertBiNode(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	ret := make([]*TreeNode, 0)
	dfs(root, &ret)
	return ret[0]
}

func dfs(node *TreeNode, ret *[]*TreeNode) {
	if node == nil {
		return
	}

	dfs(node.Left, ret)
	//这一部分要记住，
	if len(*ret) != 0 {
		node.Left = nil
		(*ret)[len(*ret) - 1].Right = node
	}
	//在判断ret不为0之后压入
	*ret = append(*ret, node)
	dfs(node.Right, ret)
}
