package mathTag

import (
	"math"
)

func FindNthDigit(n int) int {
	k := 9
	count := 1
	for n > k*count {
		n -= k*count
		k *= 10
		count++
	}
	a, b := n%count, n/count
	ret := k/9 + b - 1
	if a == 0 {
		return ret%10
	} else {
		ret += 1
	}
	ret = ret/int(math.Pow(10, float64(count-a)))
	return ret % 10
}