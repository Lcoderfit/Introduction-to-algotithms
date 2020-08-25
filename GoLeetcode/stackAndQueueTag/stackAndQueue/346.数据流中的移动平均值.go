package stackAndQueue

type Solution struct {
	queue []int
	size  int
}

func Construction(arr []int) Solution {
	return Solution{queue: arr}
}

var sum int = 0

func (this *Solution) MovingAverage(size int) {
	this.size = size
}

func (this *Solution) Next(val int) float64 {
	this.queue = append(this.queue, val)
	sum += val
	if len(this.queue) > this.size {
		this.queue = this.queue[len(this.queue)-this.size:]
	}
	return float64(sum)/float64(this.size)
}
