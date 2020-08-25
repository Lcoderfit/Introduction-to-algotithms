package tree


func TwoSumBSTs(root1 *TreeNode, root2 *TreeNode, target int) bool {
	if root1 == nil {
		return false
	}
	if find(root2, target - root1.Val) {
		return true
	}
	return TwoSumBSTs(root1.Left, root2, target) || TwoSumBSTs(root1.Right, root2, target)
}

func find(root *TreeNode, target int) bool {
	if root == nil {
		return false
	}
	if root.Val == target {
		return true
	} else if root.Val > target {
		return find(root.Left, target)
	} else {
		return find(root.Right, target)
	}
}
