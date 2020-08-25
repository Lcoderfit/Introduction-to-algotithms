package linklist


/**
* Definition for a Node.
* type Node struct {
*     Val int
*     Next *Node
*     Random *Node
* }
*/

func CopyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	p := head
	//复制节点
	for p != nil {
		s := &Node{Val: p.Val}
		s.Next = p.Next
		p.Next = s
		p = s.Next
	}
	//复制random指针
	p = head
	for p != nil {
		//仅使用一个指针p即可
		if p.Random != nil {
			p.Next.Random = p.Random.Next
		}
		p = p.Next.Next
	}
	//提取链表
	ret := head.Next
	//本题不允许修改原链表, 需要将连个链表分开
	p = head
	q := head.Next
	for p != nil {
		p.Next = q.Next
		p = q.Next
		if p != nil {
			q.Next = p.Next
			q = p.Next
		}
	}
	return ret
}