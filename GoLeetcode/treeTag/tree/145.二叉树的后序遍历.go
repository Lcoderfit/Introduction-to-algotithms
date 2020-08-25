/*
145.二叉树的后序遍历.go
时间复杂度：
空间复杂度：
*/
package tree

// 后序遍历，递归
func PostorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	left := PostorderTraversal(root.Left)
	right := PostorderTraversal(root.Right)

	ret := append(left, right...)
	ret = append(ret, root.Val)

	return ret
}

// 后序遍历，用栈
func PostorderTraversal1(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	stack := []*TreeNode{root}
	ret := make([]int, 0)
	for len(stack) != 0 {
		ret = append(ret, stack[0].Val)
		stack = stack[1:]
		if root.Right != nil {
			stack = append(stack, root.Right)
		}
		if root.Left != nil {
			stack = append(stack, root.Left)
		}
	}

	// 切片逆置
	for i, j := 0, len(ret)-1; i < j; i, j = i+1, j-1 {
		ret[i], ret[j] = ret[j], ret[i]
	}

	return ret
}
