package main

import (
	"GoLeetcode/dataStructTag/sortTag/sort"
	"fmt"
)

func main() {
	//1.merge_sort.go
	mergeSort()
}

func mergeSort() {
	arr := sort.RandomArray(10, 10)
	fmt.Println("归并排序前：", arr)
	sort.MergeSort(arr)
	fmt.Println("归并排序后：", arr)
}
