package linklist


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func DeleteNodeX(node *ListNode) {
	//不需要交换前后两个的值，直接覆盖
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}