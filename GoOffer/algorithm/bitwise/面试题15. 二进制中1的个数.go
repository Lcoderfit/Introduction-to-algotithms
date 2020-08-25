package bitwise


func HammingWeight1(n int) int {
	if n < 0 {
		n = -n
	}
	cnt := 0
	for n != 0 {
		if n & 1 == 1 {
			cnt++
		}
		n = n/2
	}
	return cnt
}

func HammingWeight(n int) int {
	if n < 0 {
		n = -n
	}
	cnt := 0
	for n != 0 {
		cnt++
		n = n&(n-1)
	}
	return cnt
}