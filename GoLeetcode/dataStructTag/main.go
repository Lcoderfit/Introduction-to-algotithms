package main

import (
	"GoLeetcode/dataStructTag/sortTag/sort"
	"GoLeetcode/dataStructTag/sortTag/sort/stable-sort"
	"fmt"
)

func main() {
	//3.merge_sort.go
	mergeSort()
}

func mergeSort() {
	arr := sort.RandomArray(10, 10)
	fmt.Println("归并排序前：", arr)
	stable_sort.MergeSort(arr)
	fmt.Println("归并排序后：", arr)
}
