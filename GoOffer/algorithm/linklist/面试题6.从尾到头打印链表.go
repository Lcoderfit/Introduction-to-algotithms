/*
6.从尾到头打印链表
时间复杂度：O(n)
空间复杂度：O(n)
*/
package linklist

// 方法一
func PrintLinkListFromTailToHead1(head *ListNode) []int {
	if head == nil {
		return []int{}
	}
	return append(PrintLinkListFromTailToHead1(head.Next), head.Val)
}

// 方法二
func PrintLinkListFromTailToHead2(head *ListNode) []int {
	if head == nil {
		return []int{}
	}

	// 反转链表
	var pre *ListNode
	for head != nil {
		pre, head, head.Next = head, head.Next, pre
	}

	ret := make([]int, 0)
	for pre != nil {
		ret = append(ret, pre.Val)
		pre = pre.Next
	}

	return ret
}

// 方法三
func PrintLinkListFromTailToHead3(head *ListNode) []int {
	if head == nil {
		return []int{}
	}

	cur := head
	n := 0
	for cur != nil {
		n++
		cur = cur.Next
	}

	ret := make([]int, n)
	cur = head
	for cur != nil {
		ret[n - 1] = cur.Val
		cur = cur.Next
		n--
	}

	return ret
}
