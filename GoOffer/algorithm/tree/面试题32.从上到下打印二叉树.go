/*
面试题32.从上到下打印二叉树.go
时间复杂度：O(n)
空间复杂度：O(n)
*/
package tree


//层序遍历：LevelOrder
func printFromTopToBottom(root *TreeNode) []int {
	ret := make([]int, 0)
	if root == nil {
		return ret
	}
	level := []*TreeNode{root}
	for len(level) != 0 {
		for _, node := range level {
			ret = append(ret, node.Val)
		}

		tmp := make([]*TreeNode, 0)
		for _, node := range level {
			tmp = append(tmp, node.Left)
			tmp = append(tmp, node.Right)
		}

		level = level[:0]
		for _, node := range tmp {
			if node != nil {
				level = append(level, node)
			}
		}
	}
	return ret
}

//递归解法
func printFromTopToBottom1(root *TreeNode, ret [][]int, index int) [][]int {
	if root == nil {
		return ret
	}
	if len(ret) < index {
		ret = append(ret, []int{})
	}
	ret[index-1] = append(ret[index-1], root.Val)
	ret = printFromTopToBottom1(root.Left, ret, index+1)
	ret = printFromTopToBottom1(root.Right, ret, index+1)
	return ret
}