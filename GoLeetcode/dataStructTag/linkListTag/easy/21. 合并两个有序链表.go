/*

方法1：
时间复杂度：O(m+n)
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
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// 一般遇到初始条件不好处理的情况时，设置一个dummy节点
	cur := &ListNode{
		Val:  -1,
		Next: nil,
	}
	head := cur

	// l1和l2控制前一个节点，cur指向当前节点，当判断出l1和l2中的更小的值时，就令cur指向对应的节点
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}

	if l1 != nil {
		cur.Next = l1
	} else {
		cur.Next = l2
	}
	return head.Next
}
