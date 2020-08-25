/*
18.删除链表的节点.go
时间复杂度：O(n)
空间复杂度：O(1)
*/
package linklist


func DeleteNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	if head.Val == val {
		return head.Next
	}

	p, q := head, head.Next
	for q != nil {
		if q.Val == val {
			p.Next = q.Next
			return head
		}
		p = q
		q = q.Next
	}

	return nil
}

// 方法二：假头
func DeleteNode1(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	extra := &ListNode{0, head}
	p, q := extra, extra.Next

	for q != nil {
		if q.Val == val {
			p.Next = q.Next
			break
		}
		p = p.Next
		q = q.Next
	}

	return extra.Next
}
