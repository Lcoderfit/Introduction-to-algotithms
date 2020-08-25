package linklist


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//注意进位处理
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	t := 0
	p, q := l1, l2
	for p.Next != nil && q.Next != nil {
		p.Val, t = (p.Val + q.Val + t) % 10, (p.Val + q.Val + t) / 10
		p = p.Next
		q = q.Next
	}
	//将q中多余的节点附在p后面
	if p.Next == nil {
		p.Next = q.Next
	}
	p.Val, t = (p.Val + q.Val + t) % 10, (p.Val + q.Val + t) / 10
	q = p
	p = p.Next
	for p != nil {
		p.Val, t = (p.Val + t) % 10, (p.Val + t) / 10
		q = p
		p = p.Next
	}
	if t > 0 {
		q.Next = &ListNode{Val: t}
	}
	return l1
}

//进阶：链表正向求和