package linklist


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	//先求链表的倒数第k+1个节点，然后再删除第k个节点
	//如果倒数第k个节点是head，那么可以通过slow == head进行判断，此时直接返回head.Next即可
	slow, fast := head, head
	//fast走了k步，fast与slow的距离为k个节点
	for ; fast != nil && n > 0; n-- {
		fast = fast.Next
	}
	if fast == nil {
		return head.Next
	}
	//需要找到倒数第k+1个节点，所以fast与slow的距离需要多一步
	fast = fast.Next
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return head
}