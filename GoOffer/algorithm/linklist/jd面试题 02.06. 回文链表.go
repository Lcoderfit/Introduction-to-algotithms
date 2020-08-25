package linklist


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//方法一：反转前半部分链表，再用快慢指针判断前后两部分是否相等
func IsPalindrome(head *ListNode) bool {
	//快慢指针，先fast一下走两步，慢指针一下走一步
	//slow指向中间（偏后）的节点，然后再令fast从head开始走
	//slow和fast每次走一步，如果slow.Val != fast.Val，则不是回文
	slow, fast := head, head
	//用于反转链表
	var pre, cur *ListNode = nil, nil
	for fast != nil && fast.Next != nil {
		cur = slow
		slow = slow.Next
		fast = fast.Next.Next
		cur.Next = pre
		pre = cur
	}

	//链表总数为奇数，需要把slow往后移动一位
	if fast != nil {
		slow = slow.Next
	}
	for slow != nil {
		if pre.Val != slow.Val {
			return false
		}
		slow = slow.Next
		pre = pre.Next
	}
	return true
}

//方法二：反转整个链表，判断反转后的链表是否与原链表相等