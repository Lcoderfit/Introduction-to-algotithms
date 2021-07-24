/*

方法1：
时间复杂度：O(n)
空间复杂度：O(1)

case1:

r:

*/
package easy

import (
	. "GoLeetcode/dataStructTag/linkListTag/common"
)

func reverseList(head *ListNode) *ListNode {
	// 如果需要控制第一个节点指向nil，则可以先设置一个pre节点初始化为nil, 然后cur指向pre
	// 就是所谓的设置dummy (假节点)
	var pre, cur, next *ListNode = nil, head, head
	for next != nil {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}
