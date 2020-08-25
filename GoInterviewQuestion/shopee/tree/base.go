package tree

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


// 层序创建二叉树
//func CreateTreeLevelOrder(valList []int) *TreeNode {
//	if valList == nil || len(valList) == 0 {
//		return nil
//	}
//	root := &TreeNode{Val: valList[0]}
//	level := []*TreeNode{root}
//	j := 1
//
//	for _, node := range level {
//		if node != nil {
//			if valList[j] != nil {
//				node.Left = &TreeNode{Val: valList[j]}
//			} else {
//				node.Left = nil
//			}
//			level = append(level, node.Left)
//			j += 1
//			if j == len(valList) {
//				return root
//			}
//
//			if valList[j] != nil {
//				node.Right = &TreeNode{Val: valList[j]}
//			} else {
//				node.Right = nil
//			}
//			level = append(level, node.Right)
//			j += 1
//			if j == len(valList) {
//				return root
//			}
//		}
//	}
//}

// 层序打印二叉树
func PrintTreeLevelOrder(root *TreeNode) {
	if root == nil {
		return
	}
	//var null *TreeNode

	ret := make([]interface{}, 0)
	level := []*TreeNode{root}
	for len(level) != 0 {
		for _, node := range level {
			if node != nil {
				ret = append(ret, node.Val)
			} else {
				ret = append(ret, nil)
			}
		}

		tmp := make([]*TreeNode, 0)
		for _, node := range level {
			if node != nil {
				tmp = append(tmp, node.Left, node.Right)
			} else {
				tmp = append(tmp, nil)
			}
		}
		copy(level, tmp)
	}
	fmt.Println("LevelOrder: ", level)
}