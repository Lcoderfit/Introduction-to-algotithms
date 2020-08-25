package linklist


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//设置a，b，c，d的意义时，应该设为a: 表示a段路径中的节点的个数(不包括入口节点)
func DetectCycleX(head *ListNode) *ListNode {
	//fast的走过的路径长需要是slow的两倍，所以初始状态为head和head.Next
	if head == nil {
		return nil
	}
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil && slow != fast {
		slow = slow.Next
		fast = fast.Next.Next
	}
	if slow != fast {
		return nil
	}
	slow = head
	//这里要注意
	fast = fast.Next
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}