/*
维护一个单调的双端队列
 */
package stackandqueue


type MaxQueue struct {
	queue []int
	maxArr []int
}


func Constructor() MaxQueue {
	return MaxQueue{queue:make([]int, 0), maxArr: make([]int, 0)}
}


func (this *MaxQueue) Max_value() int {
	if len(this.maxArr) == 0 {
		return -1
	}
	return this.maxArr[0]
}


func (this *MaxQueue) Push_back(value int)  {
	for len(this.maxArr) != 0 && value > this.maxArr[len(this.maxArr) - 1] {
		this.maxArr = this.maxArr[:len(this.maxArr) - 1]
	}
	this.maxArr = append(this.maxArr, value)
	this.queue = append(this.queue, value)
}


func (this *MaxQueue) Pop_front() int {
	if len(this.queue) == 0 {
		return -1
	}
	ret := this.queue[0]
	if ret == this.maxArr[0] {
		this.maxArr = this.maxArr[1:]
	}
	this.queue = this.queue[1:]
	return ret
}


/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */