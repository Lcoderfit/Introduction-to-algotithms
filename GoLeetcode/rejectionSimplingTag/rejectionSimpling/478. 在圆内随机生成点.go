package rejectionSimpling

import (
	"math"
	"math/rand"
	"time"
)

type Solution struct {
	rad float64
	x float64
	y float64
}


func Constructor(radius float64, x_center float64, y_center float64) Solution {
	return Solution{rad:radius, x:x_center, y:y_center}
}


func (this *Solution) RandPoint() []float64 {
	x0 := this.x - this.rad
	y0 := this.y - this.rad
	ret := make([]float64, 2)

	for {
		rand.Seed(time.Now().UnixNano())
		x := x0 + rand.Float64()*2*this.rad
		y := y0 + rand.Float64()*2*this.rad
		if math.Pow(x-this.x, 2) + math.Pow(y-this.y, 2) <= this.rad*this.rad {
			ret[0] = x
			ret[1] = y
			break
		}
	}
	return ret
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(radius, x_center, y_center);
 * param_1 := obj.RandPoint();
 */
