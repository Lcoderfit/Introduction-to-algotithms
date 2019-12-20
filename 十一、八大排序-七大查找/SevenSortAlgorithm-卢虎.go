package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("**********1.冒泡排序**********")
	array := GenRandArray(10)
	fmt.Println("排序前：", array)
	BubbleSort(array)
	fmt.Println("排序后：", array)

	fmt.Println("**********2.选择排序**********")
	array = GenRandArray(10)
	fmt.Println("排序前：", array)
	SelectionSort(array)
	fmt.Println("排序后：", array)

	fmt.Println("**********3.插入排序**********")
	array = GenRandArray(10)
	fmt.Println("排序前：", array)
	InsertionSort(array)
	fmt.Println("排序后：", array)

	fmt.Println("**********4.希尔排序**********")
	array = GenRandArray(10)
	fmt.Println("排序前：", array)
	ShellSort(array)
	fmt.Println("排序后：", array)

	fmt.Println("**********5.堆排序**********")
	array = GenRandArray(10)
	fmt.Println("排序前：", array)
	HeapSort(array)
	fmt.Println("排序后：", array)

	fmt.Println("**********6.归并排序**********")
	array = GenRandArray(10)
	fmt.Println("排序前：", array)
	MergeSort(array)
	fmt.Println("排序后：", array)

	fmt.Println("**********7.快速排序**********")
	array = GenRandArray(10)
	fmt.Println("排序前：", array)
	QuickSort(array)
	fmt.Println("排序后：", array)

}

// 生成一个包含n个元素的slice，元素取值范围为[0, n)
func GenRandArray(n int) []int {
	var array []int
	// 生成随机数种子
	rand.Seed(time.Now().Unix())
	for i := 0; i < n; i++ {
		array = append(array, rand.Intn(n))
	}
	return array
}

// 1.冒泡排序
func BubbleSort(array []int) []int {
	length := len(array)
	for i := 0; i < length-1; i++ {
		for j := 0; j < length-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	return array
}

// 2.选择排序
func SelectionSort(array []int) []int {
	length := len(array)
	for i := 0; i < length; i++ {
		for j := i; j < length; j++ {
			if array[i] > array[j] {
				array[i], array[j] = array[j], array[i]
			}
		}
	}
	return array
}

// 3.插入排序
func InsertionSort(array []int) []int {
	length := len(array)
	for index := 0; index < length; index++ {
		for (index > 0) && (array[index-1] > array[index]) {
			array[index-1], array[index] = array[index], array[index-1]
			index -= 1
		}
	}
	return array
}

// 4.希尔排序
func ShellSort(array []int) []int {
	length := len(array)
	step := length / 2

	for step > 0 {
		for cur := step; cur < length; cur++ {
			index := cur
			for (index >= step) && (array[index-step] > array[index]) {
				array[index-step], array[index] = array[index], array[index-step]
				index = index - step
			}
		}
		step = step / 2
	}
	return array
}

// 5.堆排序
func HeapSort(array []int) []int {
	first := len(array)
	for start := first; start >= 0; start-- {
		heapify(array, start, len(array)-1)
	}
	for end := len(array) - 1; end > 0; end-- {
		array[0], array[end] = array[end], array[0]
		heapify(array, 0, end-1)
	}
	return array
}

func heapify(array []int, start, end int) {
	root := start
	child := root*2 + 1
	for child <= end {
		if child+1 <= end && array[child] < array[child+1] {
			child += 1
		}
		if array[root] < array[child] {
			array[root], array[child] = array[child], array[root]
			root = child
			child = root*2 + 1
		} else {
			break
		}
	}
}

// 6.归并排序
func MergeSort(array []int) []int {
	length := len(array)
	if length > 1 {
		pivot := length / 2
		leftlen := pivot
		rightlen := length -pivot
		leftarr := make([]int, leftlen)
		rightarr := make([]int, rightlen)

		for i := 0; i < leftlen; i++ {
			leftarr[i] = array[i]
		}
		for j := 0; j < rightlen; j++ {
			rightarr[j] = array[pivot+j]
		}

		leftarr = MergeSort(leftarr)
		rightarr = MergeSort(rightarr)

		i, j, k := 0, 0, 0

		for i < leftlen && j < rightlen {
			if leftarr[i] < rightarr[j] {
				array[k] = leftarr[i]
				i += 1
			} else {
				array[k] = rightarr[j]
				j += 1
			}
			k += 1
		}

		for i < leftlen {
			array[k] = leftarr[i]
			i += 1
			k += 1
		}

		for j < rightlen {
			array[k] = rightarr[j]
			j += 1
			k += 1
		}
	}
	return array
}

// 7.快速排序
func QuickSort(array []int) []int {
	if len(array) > 1 {
		index := partition(array, 0, len(array)-1)
		QuickSort(array[:index])
		QuickSort(array[index+1:])
	}
	return array
}

func partition(array []int, low, high int) int {
	j := low - 1
	pivot := array[high]

	for i := low; i < high; i++ {
		if array[i] <= pivot {
			j += 1
			array[i], array[j] = array[j], array[i]
		}
	}

	array[j+1], array[high] = array[high], array[j+1]
	return j + 1
}
