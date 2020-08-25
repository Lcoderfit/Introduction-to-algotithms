/*
24.反转链表
时间复杂度：O(n)
空间复杂度：O(1)
*/
package linklist

// 方法一
func ReverseList1(head *ListNode) *ListNode {
	var p, q, r *ListNode = nil, head, head
	for r != nil {
		r = q.Next
		q.Next = p
		p = q
		q = r
	}
	return p
}

// 方法二
func ReverseList2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	// 会分配一块内存并返回指针, Val域初始化为0
	//pre := new(ListNode)

	var pre *ListNode
	for head != nil {
		pre, head, head.Next = head, head.Next, pre
	}

	return pre
}

