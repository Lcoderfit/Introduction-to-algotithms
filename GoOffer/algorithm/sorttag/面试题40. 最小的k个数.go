package sorttag

import (
	"container/heap"
)

//方法一：冒泡
func GetLeastNumbers(arr []int, k int) []int {
	for i := 0; i < k; i++ {
		for j := len(arr) - 1; j > i; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
	return arr[:k]
}

//方法二：堆排序
func GetLeastNumbers1(arr []int, k int) []int {
	if k <= 0 {
		return nil
	}
	if k > len(arr) {
		return arr
	}

	h := &MaxHeap{}
	*h = append(*h, arr[:k]...)
	//初始化后(*h)[0]为最大元素
	heap.Init(h)

	for i := k; i < len(arr); i++ {
		if arr[i] < (*h)[0] {
			//调用Pop方法时，会先自动将最大元素移动到slice的尾部，即(*h)[len(*h) - 1]
			heap.Pop(h)
			heap.Push(h, arr[i])
		}
	}
	return *h
}

//方法三：快排
