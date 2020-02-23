/*
35. 复杂链表的复制.go
时间复杂度：O(n)
空间复杂度：O(n)
 */
package algorithm

func CopyRandomList(head *RandomListNode) *RandomListNode {
	p := head
	// 在原来的链表的每一个节点后面先复制一个相同的节点
	for p != nil {
		newNode := &RandomListNode{p.Val, p.Number, p.Next, nil}
		p.Next = newNode
		p = newNode.Next
	}

	// 复制原节点的Random指针
	p, q := head, head.Next
	for p != nil {
		if p.Random != nil {
			q.Random = p.Random.Next
		} else {
			q.Random = nil
		}
		// 到链表最后p指向nil，p.nil不存在，所以q = p.Next会报错
		p = q.Next
		if p != nil {
			q = p.Next
		}
	}

	// 将复制的节点提取出来组合成一个新的链表
	pre, q := head.Next, head.Next
	p = pre.Next
	for q.Next != nil {
		q.Next = p.Next
		q = p.Next
		p = q.Next
	}
	return pre
}
