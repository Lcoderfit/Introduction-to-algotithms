/*
22.链表中倒数第k个节点
时间复杂度：O(k)
空间复杂度：O(1)
*/
package linklist

func GetKthFromEnd(head *ListNode, k int) *ListNode {
	low, fast := head, head
	for i := 1; i < k && fast != nil; i++ {
		fast = fast.Next
	}
	if fast == nil {
		return nil
	}

	// 因为最终fast指向最后一个节点的时候，low指向倒数第k个节点，所以判断条件为fast.Next != nil
	for fast.Next != nil {
		low = low.Next
		fast = fast.Next
	}
	return low
}
