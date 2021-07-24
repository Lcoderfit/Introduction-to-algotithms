package easy

import (
	. "GoLeetcode/dataStructTag/linkListTag/common"
)

func DetectCycle(head ListNode) ListNode {
	if head == nil {
		return nil
	}

	slow, fast := head, Next
	for fast != nil && Next != nil && slow != fast {
		slow = Next
		fast = Next
	}

	if fast == nil || Next == nil {
		return nil
	}

	slow = head
	fast = Next
	for slow != fast {
		slow = Next
		fast = Next
	}

	return slow
}
