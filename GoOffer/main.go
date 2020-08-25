package main

import (
	"GoOffer/algorithm/arrTag"
	"GoOffer/algorithm/bfsAndDfsTag"
	"GoOffer/algorithm/bitwise"
	"GoOffer/algorithm/dp"
	"GoOffer/algorithm/linklist"
	"GoOffer/algorithm/mathTag"
	"GoOffer/algorithm/sorttag"
	"GoOffer/algorithm/str"
	"GoOffer/algorithm/tree"
	"fmt"
	"runtime"
)

func main() {
	//面试题04. 二维数组中的查找.go
	//findNumberIn2DArray()

	//6.从尾到头打印链表.go
	//printListFromTailToHead()

	//10-2.青蛙跳台阶.go
	//jumpFloor()

	//面试题13. 机器人的运动范围.go
	//movingCount()

	//面试题15. 二进制中1的个数.go
	//hammingWeight()

	//18.删除链表的节点.go
	//deleteNode()

	//22.链表中倒数第k个节点.go
	//getKthFromList()

	//24.反转链表.go
	//reverseList()

	//25.合并两个排序的链表
	//mergeTwoLists()

	//面试题29. 顺时针打印矩阵.go
	//spiralOrder()

	//面试题33. 二叉搜索树的后序遍历序列.go
	//verifyPostorder()

	//35. 复杂链表的复制.go
	//copyRandomList()

	//面试题38. 字符串的排列.go
	//permutation()

	//面试题40. 最小的k个数.go
	//getLeastNumbers1()

	//面试题41. 数据流中的中位数.go
	findMedian()

	//面试题44. 数字序列中某一位的数字.go
	//findNthDigit()

	//面试题46. 把数字翻译成字符串.go
	//translateNum()

	//面试题47. 礼物的最大价值.go
	//maxValue()

	//面试题48. 最长不含重复字符的子字符串.go
	//lengthOfLongestSubstring()

	//面试题51. 数组中的逆序对.go
	//reversePairs()

	//52.两个链表的第一个公共节点.go
	//getIntersectionNode()

	//面试题56 - I. 数组中数字出现的次数.go
	//singleNumbers()
	
	//面试题57 - II. 和为s的连续正数序列.go
	//findContinuousSequence()

	//141.环形链表.go
	//hasCycle()
}

//面试题04. 二维数组中的查找.go
func findNumberIn2DArray() {

}

//6.从尾到头打印链表.go
func printListFromTailToHead() {
	array := []int{1, 2, 3, 4, 5}
	linkList, err := linklist.CreateLinkList(array)
	if err != nil {
		return
	}
	ret := linklist.PrintLinkListFromTailToHead1(linkList)
	fmt.Println("PrintListFromTailToHead1: ", ret)

	linkList, err = linklist.CreateLinkList(array)
	if err != nil {
		return
	}
	ret2 := linklist.PrintLinkListFromTailToHead2(linkList)
	fmt.Println("PrintLinkListFromTailToHead2: ", ret2)

	linkList, err = linklist.CreateLinkList(array)
	if err != nil {
		return
	}
	ret3 := linklist.PrintLinkListFromTailToHead3(linkList)
	fmt.Println("PrintLinkListFromTailToHead3: ", ret3)
}

//10-2.青蛙跳台阶.go
func jumpFloor() {
	var n int64
	var res int
	fmt.Println("/********begin**********/")
	for {
		fmt.Scanf("%d\n", &n)
		res = linklist.JumpFloor(n)
		fmt.Printf("%d级台阶跳法有：%d种\n", n, res)
	}
}

//面试题13. 机器人的运动范围.go
func movingCount() {
	var m, n, k int
	for {
		fmt.Scanf("%d%d%d", &m, &n, &k)
		ret := bfsAndDfsTag.MovingCount(m, n, k)
		fmt.Println("movingCount: ", ret)
	}
}

//面试题15. 二进制中1的个数.go
func hammingWeight() {
	var n int
	fmt.Scanf("%d\n", &n)
	for {
		ret := bitwise.HammingWeight(n)
		fmt.Println("hammingWeight: ", ret)
		fmt.Scanf("%d\n", &n)
	}
}

//18.删除链表的节点.go
func deleteNode() {
	array := []int{1, 2, 3, 4, 5}
	val := 3
	linkList, err := linklist.CreateLinkList(array)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	head := linklist.DeleteNode(linkList, val)
	linklist.PrintLinkList(head)
}

//22.链表中倒数第k个节点.go
func getKthFromList() {
	array := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	head, err := linklist.CreateLinkList(array)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	ret := linklist.GetKthFromEnd(head, k)
	fmt.Println(ret.Val)
}

//24.反转链表.go
func reverseList() {
	array := []int{1, 2, 3, 4, 5}
	head, err := linklist.CreateLinkList(array)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	ret := linklist.ReverseList1(head)
	linklist.PrintLinkList(ret)

	array = []int{5, 2, 6, 1, 9}
	head, err = linklist.CreateLinkList(array)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	ret = linklist.ReverseList2(head)
	linklist.PrintLinkList(ret)
}

//25.合并两个排序的链表.go
func mergeTwoLists() {
	valList1 := []int{1, 2, 4}
	valList2 := []int{1, 3, 4, 5}

	l1, err := linklist.CreateLinkList(valList1)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	l2, err := linklist.CreateLinkList(valList2)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	//因为是在l1和l2基础上进行连接，所以l1和l2也改变了
	//l1: 1 2 3 4 4 5
	//l2: 1 1 2 3 4 4 5
	head := linklist.MergeTwoLists1(l1, l2)
	linklist.PrintLinkList(head)

	l1, err = linklist.CreateLinkList(valList1)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	l2, err = linklist.CreateLinkList(valList2)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	head = linklist.MergeTwoLists2(l1, l2)
	linklist.PrintLinkList(head)
}

//面试题29. 顺时针打印矩阵.go
func spiralOrder() {
	var rows, cols int
	for {
		fmt.Scanf("%d%d", &rows, &cols)
		matrix := make([][]int, 0)
		for i := 0; i < rows; i++ {
			t := make([]int, cols)
			matrix = append(matrix, t)
		}
		for i := 0; i < rows; i++ {
			for j := 0; j < cols; j++ {
				var t int
				fmt.Scanf("%d", &t)
				matrix[i][j] = t
			}
		}

		ret := arrTag.SpiralOrder(matrix)
		fmt.Println("spiralOrder: ", ret)
	}
}

//面试题32.从上到下打印二叉树.go
func printFromTopToBottom() {

}

//面试题33. 二叉搜索树的后序遍历序列.go
func verifyPostorder() {
	postorder := []int{5, 2, -17, -11, 25, 76, 62, 98, 92}
	ret := tree.VerifyPostorder(postorder)
	fmt.Println("VerifyPostorder: ", ret)
}

//35. 复杂链表的复制.go
/*
	a := []interface{}{1, nil}
	for i, _ := range a {
		if a[i] == nil {
			fmt.Println(a[i])
		} else {
			fmt.Println(a[i])
		}
	}
*/
func copyRandomList() {
	valList := [][]interface{}{{7, nil}, {13, 0}, {11, 4}, {10, 2}, {1, 0}}
	head, err := linklist.CreateRandomListNode(valList)
	fmt.Println("原复杂链表：")
	linklist.PrintRandomListNode(head)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	copyHead := linklist.CopyRandomList(head)
	fmt.Println("复制后的复杂链表：")
	linklist.PrintRandomListNode(copyHead)
}

//面试题38. 字符串的排列.go
func permutation() {
	var s string
	for {
		fmt.Scanf("%s\n", &s)
		ret := str.Permutation(s)
		fmt.Println("permutation: ", ret)
	}
}

//面试题41. 数据流中的中位数.go
func findMedian() {
	var n int
	for {
		fmt.Scanf("%d", &n)
		arr := make([]int, n)
		var t int
		for i := 0; i < n; i++ {
			fmt.Scanf("%d", &t)
			arr[i] = t
		}
		ret := sorttag.FindMedian(arr)
		fmt.Println("findMedian: ", ret)
	}
}

//面试题40. 最小的k个数.go
func getLeastNumbers1() {
	arr := []int{4, 5, 1, 6, 2, 7, 3, 8}
	k := 6
	ret := sorttag.GetLeastNumbers1(arr, k)
	fmt.Println("getLeastNumbers1: ", ret)
}

//面试题44. 数字序列中某一位的数字.go
func findNthDigit() {
	var n int
	for {
		fmt.Scanf("%d", &n)
		ret := mathTag.FindNthDigit(n)
		fmt.Println("findNthDigit: ", ret)
	}
}

//面试题46. 把数字翻译成字符串.go
func translateNum() {
	var n int
	for {
		fmt.Scanf("%d", &n)
		ret := dp.TranslateNum(n)
		fmt.Println("translateNum: ", ret)
	}
}

//面试题47. 礼物的最大价值.go
func maxValue() {
	var rows, cols int
	for {
		fmt.Scanf("%d%d", &rows, &cols)
		var t int
		grid := dp.Create2DSlice(rows, cols)
		for i := 0; i < rows; i++ {
			for j := 0; j < cols; j++ {
				fmt.Scanf("%d", &t)
				grid[i][j] = t
			}
		}
		ret := dp.MaxValue(grid)
		fmt.Println("maxValue: ", ret)
	}

}

//面试题48. 最长不含重复字符的子字符串.go
func lengthOfLongestSubstring() {
	var s string
	for {
		fmt.Scanf("%s", &s)
		s = " "
		ret := str.LengthOfLongestSubstring2(s)
		fmt.Println("lengthOfLongestSubstring: ", ret)
	}
}

//面试题51. 数组中的逆序对.go
func reversePairs() {
	var n int
	for {
		fmt.Scanf("%d", &n)
		arr := make([]int, n)
		var t int
		for i := 0; i < n; i++ {
			fmt.Scanf("%d", &t)
			arr[i] = t
		}
		ret := sorttag.ReversePairs2(arr)
		fmt.Println("reversePairs: ", ret)
	}
}

//52.两个链表的第一个公共节点.go
func getIntersectionNode() {
	intersectVal := 8
	listA := []int{4, 1, 8, 4, 5}
	listB := []int{5, 0, 1, 8, 4, 5}
	skipA := 2
	skipB := 3

	headA, headB, err := linklist.CreateIntersectionLinkList(listA, listB, intersectVal, skipA, skipB)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	ret := linklist.GetIntersectionNode(headA, headB)
	fmt.Println(ret.Val)
}

//面试题56 - I. 数组中数字出现的次数.go
func singleNumbers() {
	var n int
	for {
		fmt.Scanf("%d\n", &n)
		arr := make([]int, n)
		for i := 0; i < n; i++ {
			var t int
			fmt.Scanf("%d", &t)
			arr[i] = t
		}
		ret := bitwise.SingleNumbers(arr)
		fmt.Println("singleNumbers: ", ret)
	}
}

//面试题57 - II. 和为s的连续正数序列.go
func findContinuousSequence() {
	pc, _, _, _ := runtime.Caller(0)
	fn := runtime.FuncForPC(pc).Name()

	target := 9
	ret := mathTag.FindContinuousSequence(target)
	fmt.Println(fn, ": ", ret)
}

//141.环形链表.go
func hasCycle() {

}
