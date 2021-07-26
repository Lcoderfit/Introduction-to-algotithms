/*

方法1：
时间复杂度：O()
空间复杂度：O()

case1:

r:

*/
package easy

import (
	. "GoLeetcode/dataStructTag/linkListTag/common"
)

func reverseList1(head *ListNode) *ListNode {
	// 设置一个dummy节点pre为nil，从而使得一开始的cur可以指向nil
	var pre, cur, next *ListNode = nil, head, head

	// next永远在最前面
	for next != nil {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	// pre永远在最后面，最终状态是cur和next均指向nil，而pre指向最后一个节点
	return pre
}
