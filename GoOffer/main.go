package main

import (
	"GoOffer/algorithm"
	"fmt"
)

func main() {
	// 22.链表中倒数第k个节点
	//getKthFromList()

	// 6.从尾到头打印链表
	//printListFromTailToHead()

	// 10-2.青蛙跳台阶
	//jumpFloor()

	// 24.反转链表.go
	reverseList()
}

// 24.反转链表.go
func reverseList() {
	array := []int{1, 2, 3, 4, 5}
	head, err := algorithm.CreateLinkList(array)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	ret := algorithm.ReverseList(head)
	algorithm.PrintLinkList(ret)
}

// 22.链表中倒数第k个节点
func getKthFromList() {
	array := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	head, err := algorithm.CreateLinkList(array)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	ret := algorithm.GetKthFromEnd(head, k)
	fmt.Println(ret.Val)
}

// 6.从尾到头打印链表
func printListFromTailToHead() {
	array := []int{1, 2, 3, 4, 5}
	linkList, err := algorithm.CreateLinkList(array)
	if err != nil {
		return
	}

	ret := algorithm.PrintLinkListFromTailToHead1(linkList)
	fmt.Println(ret)
}

// 10-2.青蛙跳台阶
func jumpFloor() {
	var n int64
	var res int
	fmt.Println("/********begin**********/")
	for {
		fmt.Scanf("%d\n", &n)
		res = algorithm.JumpFloor(n)
		fmt.Printf("%d级台阶跳法有：%d种\n", n, res)
	}
}