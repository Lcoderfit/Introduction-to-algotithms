package dp

import "math"

//思路清晰的代码，x,y保存商和余数
func IntegerBreak(n int) int {
	if n <= 3 {
		return n - 1
	}
	x, y := n/3, n%3
	if y == 0 {
		return int(math.Pow(3, float64(x)))
	}
	if y == 1 {
		return 4 * int(math.Pow(3, float64(x-1)))
	}
	return 2 * int(math.Pow(3, float64(x)))
}
