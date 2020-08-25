/*
面试题32 - III. 从上到下打印二叉树 III
时间复杂度：O(n)
空间复杂度：O(n)
*/
package tree


func LevelOrderIII(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	ret := make([][]int, 0)
	level := []*TreeNode{root}
	order := 1
	for len(level) != 0 {
		t1 := make([]int, 0)
		if order == 1 {
			order = 0
			for i := 0; i < len(level); i++ {
				t1 = append(t1, level[i].Val)
			}
		} else if order == 0 {
			order = 1
			for i := len(level) - 1; i >= 0; i-- {
				t1 = append(t1, level[i].Val)
			}
		}
		ret = append(ret, t1)

		t2 := make([]*TreeNode, 0)
		for _, node := range level {
			t2 = append(t2, node.Left, node.Right)
		}
		level = level[:0]
		for _, node := range t2 {
			if node != nil {
				level = append(level, node)
			}
		}
	}
	return ret
}