/*
面试题27.二叉树的镜像.go
时间复杂度：O(n)
空间复杂度：O(n) 需要一个辅助变量temp
*/
package tree

func MirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	temp := MirrorTree(root.Left)
	root.Left = MirrorTree(root.Right)
	root.Right = temp
	//可以写成下面的形式, 语法糖虽然没有辅助变量，但是背后仍然使用了一个辅助变量
	//root.Left, root.Right = MirrorTree(root.Right), MirrorTree(root.Left)

	return root
}
