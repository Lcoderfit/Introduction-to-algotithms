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
func reversePrint(head *ListNode) []int {
	if head == nil {
		return nil
	}
	ret := make([]int, 0)
	p := head
	for p != nil {
		ret = append(ret, p.Val)
		p = p.Next
	}

	n := len(ret)
	for i := 0; i < len(ret)/2; i++ {
		ret[i], ret[n-i-1] = ret[n-i-1], ret[i]
	}
	return ret
}
