/*
52.两个链表的第一个公共节点.go
时间复杂度：O(n)
空间复杂度：O(1)
*/
package linklist


func GetIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	p, q := headA, headB
	for p != q {
		// 就算两个链表没有交点，将nil作为一次状态，到最后双指针都会同时指向nil
		if p == nil {
			p = headB
		} else {
			p = p.Next
		}

		if q == nil {
			q = headA
		} else {
			q = q.Next
		}
	}

	return p
}
