package stackandqueue


type CQueue struct {
	stackA []int
	stackB []int
}


func Constructor() CQueue {
	return CQueue{
		stackA: make([]int, 0),
		stackB: make([]int, 0),
	}
}


func (this *CQueue) AppendTail(value int)  {
	this.stackA = append(this.stackA, value)
}


func (this *CQueue) DeleteHead() int {
	if len(this.stackB) == 0 {
		for len(this.stackA) != 0 {
			this.stackB = append(this.stackB, this.stackA[len(this.stackA) - 1])
			this.stackA = this.stackA[:len(this.stackA) - 1]
		}
	}

	if len(this.stackB) == 0 {
		return -1
	}
	ret := this.stackB[len(this.stackB) - 1]
	this.stackB = this.stackB[:len(this.stackB) - 1]
	return ret
}


/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
