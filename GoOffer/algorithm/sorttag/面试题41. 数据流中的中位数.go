package sorttag

import "container/heap"

//方法二：优先队列，用两个堆，小根堆存较大的一半，大根堆存较小的一半
//设置最大堆和最小堆
type MedianFinder struct {
	//小根堆存贮较大的一般，大根堆存贮较小的一般
	//如果总数为奇数，放小根堆中
	minHeap *Heap
	maxHeap *Heap
}


/** initialize your data structure here. */
func Constructor() MedianFinder {
	//初始化大根堆和小根堆,必须为指针类型
	h1, h2 := &Heap{order:true}, &Heap{order:false}
	//container/heap库函数进行初始化
	heap.Init(h1)
	heap.Init(h2)
	return MedianFinder{
		minHeap: h1,
		maxHeap: h2,
	}
}


func (this *MedianFinder) AddNum(num int)  {
	if this.minHeap.Len() > this.maxHeap.Len() {
		//先放入小根堆，再放入大根堆
		// this.minHeap.Push(num)
		// this.maxHeap.Push(this.minHeap.Pop())
		heap.Push(this.minHeap, num)
		heap.Push(this.maxHeap, heap.Pop(this.minHeap))
	} else {
		//先放入大根堆中，再放入小根堆
		// this.maxHeap.Push(num)
		// this.minHeap.Push(this.maxHeap.Pop())
		heap.Push(this.maxHeap, num)
		heap.Push(this.minHeap, heap.Pop(this.maxHeap))
	}
}


func (this *MedianFinder) FindMedian() float64 {
	if this.minHeap.Len() == 0 {
		return 0
	}
	if this.minHeap.Len() > this.maxHeap.Len() {
		return float64(this.minHeap.arr[0])
	} else {
		return float64(this.minHeap.arr[0] + this.maxHeap.arr[0])/float64(2)
	}
}


/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */