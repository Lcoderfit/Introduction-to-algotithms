package algorithm

func ReverseList(head *ListNode) *ListNode {
	var p, q, r *ListNode = nil, head, head
	for r != nil {
		r = q.Next
		q.Next = p
		p = q
		q = r
	}
	return p
}
