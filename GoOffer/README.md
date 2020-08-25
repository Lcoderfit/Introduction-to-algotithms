## 3.1

1.Go语言数据结构的使用（list/queue等模块）



2.判断切片等是否不含有元素时，用for len(a) != 0, Python用 if not a(Python默认所有的不含元素的数据结构都为False)

3.做树一类的算法采用递归时，用单点二倒推法，先假设只有一个节点的情况（if root == nil {return nil}），然后假设有两层，一个根节点，左右各有一个节点，根节点获取左子节点和右子节点递归返回的值，然后与根节点合并



4.不定参：

func a(val ...int)点在前

append(b, []int{1, 2}...)点在后



5.初始化二维slice：a := `[][]int{{1}}`, 只有单个元素时可以不加逗号

6.将slice B复制到A，copy(A, B), A的长度必须大于B

7.递归和回溯比较相似，递归是先判断特殊情况，然后先返回，再获取递归结果，最后返回整合结果

```
func IsSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}

	res := dfs(A, B)
	if !res {
		res = IsSubStructure(A.Left, B)
	}
	if !res {
		res = IsSubStructure(A.Right, B)
	}

	//利用 || 的短路特性可写成
	//return dfs(A,B) || isSubStructure(A.Left,B) || isSubStructure(A.Right,B)
	return res
}

func dfs(a, b *TreeNode) bool {
	if b == nil {
		return true
	}
	if a == nil {
		return false
	}

	if a.Val != b.Val {
		return false
	}
	return dfs(a.Right, b.Right) && dfs(a.Left, b.Left)
}
```



回溯是满足情况先递归，递归完后，不满足情况再向上回溯一步

```
func dfs(root *TreeNode, sum int, arr []int, ret *[][]int) {
	arr = append(arr, root.Val)

	if root.Val == sum && root.Left == nil && root.Right == nil {
		//slice是一个指向底层的数组的指针结构体
		//因为是先序遍历，如果 root.Right != nil ,arr 切片底层的数组会被修改
		//所以这里需要 copy arr 到 tmp，再添加进 ret，防止 arr 底层数据修改带来的错误
		tmp := make([]int, len(arr))
		copy(tmp, arr)
		*ret = append(*ret, tmp)
	}

	if root.Left != nil {
		dfs(root.Left, sum-root.Val, arr, ret)
	}
	if root.Right != nil {
		dfs(root.Right, sum-root.Val, arr, ret)
	}
	arr = arr[:len(arr)-1] //回溯
}
```



