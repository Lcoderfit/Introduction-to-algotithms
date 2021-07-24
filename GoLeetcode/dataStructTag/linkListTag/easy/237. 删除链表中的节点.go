/*

方法1：
时间复杂度：O(1)
空间复杂度：O(1)

case1:

r:

*/
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
func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
