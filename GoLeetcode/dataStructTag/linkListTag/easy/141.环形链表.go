/*
141.环形链表.go
时间复杂度：O(n)
空间复杂度：O(1)
*/
package easy

import (
	"GoLeetcode/dataStructTag/linkListTag/common"
)

func HasCycle(head *common.ListNode) bool {
	if head == nil {
		return false
	}

	slow, fast := head, Next
	for fast != nil && Next != nil {
		if slow == fast {
			return true
		}
		slow = Next
		fast = Next
	}

	return false
}
