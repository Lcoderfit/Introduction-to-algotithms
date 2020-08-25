package linklist


//方法一：快慢指针
func KthToLast(head *ListNode, k int) int {
	//快慢指针
	fast, slow := head, head
	for ; fast != nil && k > 1 ; k-- {
		fast = fast.Next
	}
	if fast == nil {
		return -1
	}
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	return slow.Val
}

//二：快慢指针
func kthToLast(head *ListNode, k int) int {
	//快慢指针
	fast, slow := head, head
	//快指针走了k步，到了第k+1个节点处
	//所以判断后面快慢指针一起走时的判断条件是fast != nil
	for ; fast != nil && k > 0; k-- {
		fast = fast.Next
	}
	if k > 0 {
		return -1
	}
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	return slow.Val
}