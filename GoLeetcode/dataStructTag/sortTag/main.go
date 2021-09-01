package main

import (
	"GoLeetcode/dataStructTag/sortTag/sort/stable-sort"
	"GoLeetcode/dataStructTag/sortTag/sort/unstable-sort"
	"fmt"
)

func main() {
	// 1.冒泡排序
	// bubbleSort()
	// bubbleSort1()

	// 2.插入排序
	//insertSort()

	// 3.归并排序
	mergeSort()

	// 4.selectSort
	//selectSort()

	// 5.quickSort
	//quickSort()

	// 6.堆排序
	//heapSort()
}

// 1.冒泡排序
func bubbleSort() {
	var n int
	for {
		fmt.Scanln(&n)
		arr := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Scanf("%d", &arr[i])
		}
		res := stable_sort.BubbleSort(arr)
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
		res := stable_sort.BubbleSort1(arr)
		fmt.Println(res)
	}
}

// 2.插入排序
func insertSort() {
	var n int
	for {
		fmt.Scanln(&n)
		arr := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Scanf("%d", &arr[i])
		}
		res := stable_sort.InsertSort(arr)
		fmt.Println(res)
	}
}

// 3.归并排序
func mergeSort() {
	var n int
	for {
		fmt.Scanln(&n)
		arr := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Scanf("%d", &arr[i])
		}
		res := stable_sort.MergeSort(arr)
		fmt.Println(res)
	}
}

// 4.选择排序
func selectSort() {
	var n int
	for {
		fmt.Scanln(&n)
		arr := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Scanf("%d", &arr[i])
		}
		res := unstable_sort.SelectSort(arr)
		fmt.Println(res)
	}
}

// 5.快速排序
func quickSort() {
	var n int
	for {
		fmt.Scanln(&n)
		arr := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Scanf("%d", &arr[i])
		}
		unstable_sort.QuickSort(arr, 0, len(arr)-1)
		fmt.Println(arr)
	}
}

// 6.堆排序
func heapSort() {
	var n int
	for {
		fmt.Scanln(&n)
		arr := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Scanf("%d", &arr[i])
		}
		res := unstable_sort.HeapSort(arr)
		fmt.Println(res)
	}
}
