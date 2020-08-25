package str

//最大公约数
func MaxGcd1(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func MaxGcd2(x, y int) int {
	for y != 0 {
		if x > y {
			x = x - y
		} else {
			y = y - x
		}
	}
	return x
}

//最小公倍数
func MinGcd1(x, y int) int {
	s := x * y
	for y != 0 {
		x, y = y, x % y
	}
	return s/x
}

func MinGcd2(x, y int) int {
	s := x * y
	for y != 0 {
		if x > y {
			x = x - y
		} else {
			y = y - x
		}
	}
	return s/x
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}