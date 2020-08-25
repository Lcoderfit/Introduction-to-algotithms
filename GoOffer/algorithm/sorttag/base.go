package sorttag

import (
	"container/heap"
	"fmt"
)

func UseHeapExample(arr []int) []int {
	h := &Heap{}
	h.order = false
	h.arr = append(h.arr, arr...)
	//初始化后(*h)[0]为最大元素
	heap.Init(h)

	m := &Heap{}
	m.order = true
	m.arr = append(m.arr, arr...)
	heap.Init(m)
	fmt.Println("min: ", m)
	return h.arr
}

//添加一个属性order，order为true,则是最小堆，order为false，就是最大堆
type Heap struct {
	arr []int
	order bool
}

func (h *Heap) Len() int {
	return len(h.arr)
}

func (h *Heap) Less(i, j int) bool {
	//最小堆
	if h.order {
		return h.arr[i] < h.arr[j]
	} else {
		//大根堆
		return h.arr[i] > h.arr[j]
	}
}

func (h *Heap) Swap(i, j int) {
	h.arr[i], h.arr[j] = h.arr[j], h.arr[i]
}

//注意，这里就这么写，到时候调用时调用的是heap库的Push方法，而不是h.Push()
//如果要调用h.Push(), 则需要修改实现，实现堆排序的逻辑
func (h *Heap) Push(x interface{}) {
	h.arr = append(h.arr, x.(int))
}

func (h *Heap) Pop() interface{} {
	x := h.arr[len(h.arr) - 1]
	h.arr = h.arr[:len(h.arr) - 1]
	return x
}


//最小堆
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
//大根堆, 注意，根在slice的尾部，而不是首部，例如：5, 4, 1, 6; 但是用fmt.Println(*h)输出时，输出的是6, 5, 1, 4
//如果是小根堆改成h[i] < h[j]
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}


//最大堆
type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
//大根堆, 注意，根在slice的尾部，而不是首部，例如：5, 4, 1, 6; 但是用fmt.Println(*h)输出时，输出的是6, 5, 1, 4
//如果是小根堆改成h[i] < h[j]
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}