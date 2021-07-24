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
//func deleteDuplicates(head *ListNode) *ListNode {
//	if head == nil {
//		return nil
//	}
//	pre, cur := head, head.Next
//	for cur != nil {
//		if pre.Val != cur.Val {
//			pre.Next = cur
//			pre = cur
//		}
//		cur = cur.Next
//	}
//	return head
//}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	cur := head
	// cur.Next != nil 与 cur.Next = cur.Next.Next连用
	for cur.Next != nil {
		if cur.Val == cur.Next.Val {
			// 用一个指针删除cur的后一个节点
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}
