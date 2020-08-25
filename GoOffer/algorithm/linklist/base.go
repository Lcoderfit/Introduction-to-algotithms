package linklist

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func CreateLinkList(valList []int) (*ListNode, error) {
	if len(valList) <= 0 {
		return nil, fmt.Errorf("valList is empty")
	}
	head := &ListNode{valList[0], nil}
	p := head
	for i := 1; i < len(valList); i++ {
		p.Next = &ListNode{valList[i], nil}
		p = p.Next
	}
	return head, nil
}

func PrintLinkList(head *ListNode) {
	if head == nil {
		return
	}
	p := head
	for p != nil {
		fmt.Printf("%d ", p.Val)
		p = p.Next
	}
	fmt.Println()
}

// 复杂链表
type RandomListNode struct {
	Val int
	// 存储节点的random指针所指向的节点的序号, 有可能为nil，所以是interface{}
	Number interface{}
	Next *RandomListNode
	Random *RandomListNode
}

// 创建RandomListNode
func CreateRandomListNode(valList [][]interface{}) (*RandomListNode, error) {
	if len(valList) <= 0 {
		return nil, fmt.Errorf("valList is empty")
	}

	// 通过number找到对应的节点地址
	randomMap := map[int]*RandomListNode{}
	head := &RandomListNode{valList[0][0].(int), nil, nil, nil}
	randomMap[0] = head
	// 先进行Next的连接
	p := head
	for i := 1; i < len(valList); i++ {
		p.Next = &RandomListNode{valList[i][0].(int), nil, nil, nil}
		p = p.Next
		// 存储自身节点序号与自身指针的映射
		randomMap[i] = p
	}

	// 进行Random的连接
	p = head
	for i := 0; p != nil; i++ {
		if valList[i][1] != nil {
			p.Number = valList[i][1].(int)
			p.Random = randomMap[p.Number.(int)]
		} else {
			p.Number = nil
			p.Random = nil
		}
		p = p.Next
	}

	return head, nil
}

// 打印RandomListNode
func PrintRandomListNode(head *RandomListNode) {
	// 定义一个双重切片
	ret := make([][]interface{}, 0)
	p := head
	for p != nil {
		ret = append(ret, []interface{}{p.Val, p.Number})
		p = p.Next
	}
	fmt.Println(ret)
}

// 创建Y形链表，两个链表有一个公共节点
func CreateIntersectionLinkList(listA, listB []int, intersectVal, skipA, skipB int) (*ListNode, *ListNode, error) {
	if skipA > len(listA) || skipB > len(listB) || (len(listA) - skipA != len(listB) - skipB) {
		return nil, nil, fmt.Errorf("listA or listB is invalid")
	}
	if (skipA == 0 && skipB != 0) || (skipB == 0 && skipA != 0) {
		return nil, nil, fmt.Errorf("skipA or skipB value error")
	}
	if (skipA == len(listA) && skipB != len(listB)) || (skipB == len(listB) && skipA != len(listA)) {
		return nil, nil, fmt.Errorf("skipA or skipB not match")
	}
	//if listA[skipA] != intersectVal || listB[skipB] != intersectVal {
	//	return nil, nil, fmt.Errorf("intersecterVal is invalid")
	//}
	//for i := skipA; i < len(listA); i++ {
	//	j := skipB
	//	if listA[i] != listB[j] {
	//		return nil, nil, fmt.Errorf("listA or listB cant't construct IntersectionLinkList")
	//	}
	//	j++
	//}

	//headA := &ListNode{listA[0], nil}
	//p := headA
	//intersectNode := headA
	//for i := 1; i < len(listA); i++ {
	//	p.Next = &ListNode{listA[i], nil}
	//	p = p.Next
	//	if skipA == i {
	//		intersectNode = p
	//	}
	//}
	//if intersectNode == headA {
	//	return headA, headA, nil
	//}
	//
	//headB := &ListNode{listB[0], nil}
	//q := headB
	//for i := 1; i < skipB; i++ {
	//	q.Next = &ListNode{listB[i], nil}
	//	q = q.Next
	//}
	//
	//q.Next = intersectNode
	//
	//return headA, headB, nil

	headA := &ListNode{listA[0], nil}
	p := headA
	if skipA == 0 {
		for i := 1; i < len(listA); i++ {
			p.Next = &ListNode{listA[i], nil}
			p = p.Next
		}
		return headA, headA, nil
	} else if skipA == len(listA) {
		for i := 1; i < len(listA); i++ {
			p.Next = &ListNode{listA[i], nil}
			p = p.Next
		}
		headB := &ListNode{listB[0], nil}
		q := headB
		for i := 1; i < len(listB); i++ {
			q.Next = &ListNode{listB[i], nil}
			q = q.Next
		}
		return headA, headB, nil
	} else {
		var intersectNode *ListNode
		for i := 1; i < len(listA); i++ {
			p.Next = &ListNode{listA[i], nil}
			p = p.Next
			if skipA == i {
				intersectNode = p
			}
		}

		headB := &ListNode{listB[0], nil}
		q := headB
		for i := 1; i < skipB; i++ {
			q.Next = &ListNode{listB[i], nil}
			q = q.Next
		}
		q.Next = intersectNode
		return headA, headB, nil
	}
}