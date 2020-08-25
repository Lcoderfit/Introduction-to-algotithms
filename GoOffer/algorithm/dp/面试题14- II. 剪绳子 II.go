package dp


func CuttingRopeII(n int) int {
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}
	t3, ret := n/3, 1
	if n % 3 == 1 {
		ret = 4
		t3 -= 1
	}
	if n % 3 == 2 {
		ret = 2
	}
	for i := 0; i < t3; i++ {
		ret = ret*3 % (1e9+7)
	}
	return ret
}