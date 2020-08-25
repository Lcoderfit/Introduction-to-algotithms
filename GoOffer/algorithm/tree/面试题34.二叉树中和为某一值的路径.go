/*
面试题34.二叉树中和为某一值的路径.go
时间复杂度：O(n)???
空间复杂度：O(n)???
*/
package tree

func PathSum(root *TreeNode, sum int) [][]int {
	ret := make([][]int, 0)
	if root == nil {
		return ret
	}

	if root.Left == nil && root.Right == nil && root.Val == sum {
		return [][]int{{root.Val}}
	}
	left := PathSum(root.Left, sum-root.Val)
	right := PathSum(root.Right, sum-root.Val)

	for _, item := range append(left, right...) {
		tmp := append([]int{root.Val}, item...)
		ret = append(ret, tmp)
	}

	return ret
}

//func pathSum(root *TreeNode, sum int) [][]int {
//	if root == nil {
//		return nil
//	}
//	var ret [][]int
//	dfs(root, sum, []int{}, &ret)
//	return ret
//}
//
//func dfs(root *TreeNode, sum int, arr []int, ret *[][]int) {
//	arr = append(arr, root.Val)
//
//	if root.Val == sum && root.Left == nil && root.Right == nil {
//		//slice是一个指向底层的数组的指针结构体
//		//因为是先序遍历，如果 root.Right != nil ,arr 切片底层的数组会被修改
//		//所以这里需要 copy arr 到 tmp，再添加进 ret，防止 arr 底层数据修改带来的错误
//		tmp := make([]int, len(arr))
//		copy(tmp, arr)
//		*ret = append(*ret, tmp)
//	}
//
//	if root.Left != nil {
//		dfs(root.Left, sum-root.Val, arr, ret)
//	}
//	if root.Right != nil {
//		dfs(root.Right, sum-root.Val, arr, ret)
//	}
//	arr = arr[:len(arr)-1] //回溯
//}
