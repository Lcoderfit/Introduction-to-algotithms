package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func LevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	ret := make([][]int, 0)
	//存储每一层的节点（不包含空节点）
	level := []*TreeNode{root}
	for len(level) != 0 {
		levelNodeVal := make([]int, 0)
		for _, node := range level {
			if node != nil {
				levelNodeVal = append(levelNodeVal, node.Val)
			}
		}
		ret = append(ret, levelNodeVal)

		//存储每一层的节点，可能包含空节点
		tmp := make([]*TreeNode, 0)
		for _, node := range level {
			tmp = append(tmp, node.Left, node.Right)
		}

		//清空
		level = level[:0]
		for _, node := range tmp {
			if node != nil {
				level = append(level, node)
			}
		}
	}
	return ret
}

//递归解法，调用方法：LevelOrder(root, ret, 1)
func LevelOrder1(root *TreeNode, ret [][]int, index int) [][]int {
	if root == nil {
		return ret
	}
	//遍历到更深的一层，再ret中添加一个slice
	if len(ret) < index {
		ret = append(ret, []int{})
	}

	//把当前层的节点的Val存入ret中
	ret[index-1] = append(ret[index-1], root.Val)
	LevelOrder1(root.Left, ret, index+1)
	LevelOrder1(root.Right, ret, index+1)
}
