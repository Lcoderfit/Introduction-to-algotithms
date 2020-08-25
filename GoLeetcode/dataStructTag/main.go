package main

import (
	"GoLeetcode/dataStructTag/sortTag"
	"fmt"
)

func main() {
	//1.merge_sort.go
	mergeSort()
}

func mergeSort() {
	arr := sortTag.RandomArray(10, 10)
	fmt.Println("归并排序前：", arr)
	sortTag.MergeSort(arr)
	fmt.Println("归并排序后：", arr)
}