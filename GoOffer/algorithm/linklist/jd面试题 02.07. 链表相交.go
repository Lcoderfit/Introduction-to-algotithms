package linklist


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//两个链表相交是Y形的，不是X形的
func GetIntersectionNodeX(headA, headB *ListNode) *ListNode {
	p, q := headA, headB
	for p != q {
		// 就算两个链表没有交点，将nil作为一次状态，到最后双指针都会同时指向nil
		//这样写不会跳过头节点, 把末尾的null也算作跳跃的一步
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