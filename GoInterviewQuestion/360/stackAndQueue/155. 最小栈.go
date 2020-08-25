package stackAndQueue


type MinStack struct {
	data []int
	min []int
}


/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		data: make([]int, 0),
		min: make([]int, 0),
	}
}


func (this *MinStack) Push(x int)  {
	this.data = append(this.data, x)
	//控制s2非严格单调递减，这样末尾的元素总是对应当前s1中最小的元素
	if len(this.min) == 0 || x <= this.min[len(this.min) - 1] {
		this.min = append(this.min, x)
	}
}


func (this *MinStack) Pop()  {
	if len(this.data) == 0 {
		return
	}
	x := this.data[len(this.data) - 1]
	this.data = this.data[:len(this.data) - 1]
	if x == this.min[len(this.min) - 1] {
		this.min = this.min[:len(this.min) - 1]
	}
}


func (this *MinStack) Top() int {
	if len(this.data) == 0 {
		return -1
	}
	return this.data[len(this.data) - 1]
}


func (this *MinStack) GetMin() int {
	if len(this.min) == 0 {
		return -1
	}
	return this.min[len(this.min) - 1]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */