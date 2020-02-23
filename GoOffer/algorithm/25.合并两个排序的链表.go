/*
25.合并两个排序的链表
时间复杂度：O(n)
空间复杂度：O(n)
*/
package algorithm

// 迭代解法
func MergeTwoLists1(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil{
		return l1
	}

	var head *ListNode
	if l1.Val < l2.Val {
		head = l1
		l1 = l1.Next
	} else {
		head = l2
		l2 = l2.Next
	}

	p := head
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			p.Next = l1
			l1 = l1.Next
		} else {
			p.Next = l2
			l2 = l2.Next
		}
		p = p.Next
	}
	if l1 != nil {
		p.Next = l1
	} else if l2 != nil {
		p.Next = l2
	}
	return head
}

// 递归解法
func MergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	if l1.Val < l2.Val {
		l1.Next = MergeTwoLists2(l1.Next, l2)
		return l1
	}

	l2.Next = MergeTwoLists2(l1, l2.Next)
	return l2
}