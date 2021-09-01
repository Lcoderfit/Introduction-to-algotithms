package main

import (
	"GoLeetcode/dataStructTag/sortTag/sort"
	"fmt"
)

func main() {
	//bubbleSort()
	//bubbleSort1()
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(len(arr), cap(arr))

	arr1 := arr[:]
	arr1[0] = 10
	fmt.Println(len(arr1), cap(arr1))
	fmt.Println(arr, arr1)

	arr2 := arr[0:0]
	fmt.Println(len(arr2), cap(arr2))
	arr2 = append(arr2, 8)
	fmt.Println(arr2, arr)

}

func bubbleSort() {
	var n int
	for {
		fmt.Scanln(&n)
		arr := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Scanf("%d", &arr[i])
		}
		res := sort.BubbleSort(arr)
		fmt.Println(res)
	}
}

func bubbleSort1() {
	var n int
	for {
		fmt.Scanln(&n)
		arr := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Scanf("%d", &arr[i])
		}
		res := sort.BubbleSort1(arr)
		fmt.Println(res)
	}
}
