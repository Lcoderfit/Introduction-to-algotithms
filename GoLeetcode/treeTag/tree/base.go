package tree


type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func CreateTree(valList []int) *Node {

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}