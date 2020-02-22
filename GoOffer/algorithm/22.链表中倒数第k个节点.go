/*
22.链表中倒数第k个节点
时间复杂度：O(1)
空间复杂度：O(1)
*/
package algorithm

func GetKthFromEnd(head *ListNode, k int) *ListNode {
	low, fast := head, head
	for i := 1; i < k && fast != nil; i++ {
		fast = fast.Next
	}
	if fast == nil {
		return nil
	}

	for fast.Next != nil {
		low = low.Next
		fast = fast.Next
	}
	return low
}
