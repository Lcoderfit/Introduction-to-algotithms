/*
方法1：动态规划
时间复杂度：O(n) n表示数的节点总数
空间复杂度：O(n)

case1:
3 2 3 null 3 null 1
r:
7

case2:
3 4 5 1 3 null 1
r:
9
*/
package medium

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Robe 能偷到的最高金额
func Robe(root *TreeNode) int {
	ans := dfs(root)
	return Max(ans[0], ans[1])
}

// dfs 找到取根节点和不取根节点时的满足条件的最大值；返回一个切片，
// 切片只有两个元素，第一个元素是选择根节点的情况下所能获取的最大值，第二个元素是不选择根节点的情况下所能获取的最大值
func dfs(node *TreeNode) []int {
	if node == nil {
		return []int{0, 0}
	}
	l, r := dfs(node.Left), dfs(node.Right)
	selected := node.Val + l[1] + r[1]
	notSelected := Max(l[0], l[1]) + Max(r[0], r[1])
	return []int{selected, notSelected}
}
