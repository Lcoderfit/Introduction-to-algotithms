/*

方法1：
时间复杂度：O(n)
空间复杂度：O(1)

case1:

r:

*/
package easy

import (
	. "GoLeetcode/dataStructTag/linkListTag/common"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getKthFromEnd(head *ListNode, k int) *ListNode {
	// 对付倒数第k，设置dummy并令cur走到nil结束
	var pre, cur *ListNode = nil, head
	i := 0
	// 向前走k步，即cur位于第k+1个节点
	for i < k && cur != nil {
		pre = cur
		cur = cur.Next
		i++
	}
	if i != k {
		return nil
	}

	// 这一步容易漏写，需要重新设置pre为head节点,此时pre与cur中间间隔k-1个节点（不包括pre和cur所在节点）
	// 当cur到达nil时，pre仍与cur间隔k-1个节点, k-1个节点加上pre自己本身，即pre位于倒数第k个节点
	pre = head
	for cur != nil {
		cur = cur.Next
		pre = pre.Next
	}
	return pre
}
