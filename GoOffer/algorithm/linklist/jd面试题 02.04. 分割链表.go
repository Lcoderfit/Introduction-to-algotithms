package linklist


func Partition(head *ListNode, x int) *ListNode {
	p, q := head, head
	//快排思想解决
	for ;q != nil; q = q.Next {
		if q.Val < x {
			p.Val, q.Val = q.Val, p.Val
			p = p.Next
		}
	}
	return head
}