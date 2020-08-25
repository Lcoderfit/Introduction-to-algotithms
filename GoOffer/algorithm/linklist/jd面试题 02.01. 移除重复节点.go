package linklist


//方法二：添加一个假头
func RemoveDuplicateNodes(head *ListNode) *ListNode {
	//根据长度和大小限制，可以打个表
	dummy := &ListNode{Next: head}
	m := make(map[int]int, 0)
	for cur := dummy; cur.Next != nil; {
		if m[cur.Next.Val] != 0 {
			cur.Next = cur.Next.Next
		} else {
			m[cur.Next.Val]++
			cur = cur.Next
		}
	}
	return head
}