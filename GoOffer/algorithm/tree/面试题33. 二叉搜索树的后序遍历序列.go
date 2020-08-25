/*
面试题33. 二叉搜索树的后序遍历序列.go
时间复杂度：O(n^2): 最坏, n+n-1+n-2+n-3...
空间复杂度：O(n)
*/
package tree

func VerifyPostorder(postorder []int) bool {
	if postorder == nil {
		return false
	}
	if len(postorder) <= 2 {
		return true
	}

	return isBinary(postorder)
}

func isBinary(array []int) bool {
	if len(array) <= 1 {
		return true
	}
	root := array[len(array)-1]
	var pos int
	for pos < len(array)-1 && array[pos] < root {
		pos++
	}

	for i := pos; i < len(array)-1; i++ {
		if array[i] < root {
			return false
		}
	}
	return isBinary(array[:pos]) && isBinary(array[pos:len(array)-1])
}

//非递归解法
//func VerifyPostorder1(postorder []int) bool {
//
//}