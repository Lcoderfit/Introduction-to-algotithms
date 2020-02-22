/*
6.从尾到头打印链表
时间复杂度：O(n)
空间复杂度：O(n)
*/
package algorithm

// 方法一
func PrintLinkListFromTailToHead1(head *ListNode) []int {
	if head == nil {
		return []int{}
	}
	return append(PrintLinkListFromTailToHead1(head.Next), head.Val)
}

// 方法二
//func PrintLinkListFormTailToHead2()