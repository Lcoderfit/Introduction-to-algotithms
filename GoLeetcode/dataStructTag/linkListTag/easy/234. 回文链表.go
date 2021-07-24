package easy

import (
	. "GoLeetcode/dataStructTag/linkListTag/common"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return false
	}
	ret := make([]int, 0)
	p := head
	for p != nil {
		ret = append(ret, p.Val)
		p = p.Next
	}

	n := len(ret)
	for i := 0; i < len(ret)/2; i++ {
		if ret[i] != ret[n-i-1] {
			return false
		}
	}
	return true
}
