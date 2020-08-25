/*
面试题32 - II. 从上到下打印二叉树 II.go
时间复杂度：O(n)
空间复杂度：O(n)
*/
package tree


func LevelOrderII(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	level := []*TreeNode{root}
	ret := make([][]int, 0)
	for len(level) != 0 {
		a := make([]int, 0)
		for _, node := range level {
			a = append(a, node.Val)
		}
		ret = append(ret, a)

		tmp := make([]*TreeNode, 0)
		for _, node := range level {
			tmp = append(tmp, node.Left, node.Right)
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