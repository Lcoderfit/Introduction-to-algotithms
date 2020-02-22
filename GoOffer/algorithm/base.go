package algorithm

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
