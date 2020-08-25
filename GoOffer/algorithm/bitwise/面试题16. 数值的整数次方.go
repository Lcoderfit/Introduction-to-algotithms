package bitwise


func MyPow(x float64, n int) float64 {
	if n < 0 {
		x, n = 1/x, -n
	}
	ret := 1.0
	for n > 0 {
		if n & 1 == 1 {
			ret *= x
		}
		x = x * x
		n /= 2
	}
	return ret
}