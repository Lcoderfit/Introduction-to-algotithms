package tree


func KthLargest(root *TreeNode, k int) int {
	if root == nil {
		return 0
	}

	count := 0
	var ret *TreeNode
	helper(root, &count, k, &ret)
	if ret == nil {
		return 0
	}
	return ret.Val
}

func helper(node *TreeNode, count *int, k int, ret **TreeNode) {
	if node == nil {
		return
	}
	if node.Right != nil {
		helper(node.Right, count, k, ret)
	}
	*count += 1
	if k == *count {
		*ret = node
		return
	}
	if node.Left != nil {
		helper(node.Left, count, k, ret)
	}
	return
}

//方法二
//var count int
//var res int
//
//func KthLargest1(root *TreeNode, k int) int {
//	count = k
//	helper1(root)
//	return res
//}
//
//func helper1(root *TreeNode){
//	if root == nil {
//		return
//	}
//	helper1(root.Right)
//
//	if count == 1 {
//		res = root.Val
//		count--
//		return
//	}
//
//	count--
//	helper1(root.Left)
//}

var count int
var res int

func KthLargest1(root *TreeNode, k int) int {
	count = k
	helper1(root)
	return res
}

func helper1(node *TreeNode) {
	if node == nil {
		return
	}

	helper1(node.Left)

	if count == 1 {
		res = node.Val
		// 必须加上这一步
		count--
		return
	}

	count--
	helper1(node.Right)
}

