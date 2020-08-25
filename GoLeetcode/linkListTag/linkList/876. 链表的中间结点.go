package linkList


func MiddleNode(head *ListNode) *ListNode {
	//快慢指针，fast和low从起始点同时出发
	//low一次走一步，fast一次走两步
	//如果fast.next == nil,直接返回low
	//如果fast == nil, 返回low
	low, fast := head, head
	for fast != nil && fast.Next != nil {
		low = low.Next
		fast = fast.Next.Next
	}
	return low
}