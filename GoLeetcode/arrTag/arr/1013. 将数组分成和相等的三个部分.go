package arr


//方法一
func CanThreePartsEqualSum(A []int) bool {
	num := 0
	for _, v := range A {
		num += v
	}
	if num % 3 != 0 {
		return false
	}
	target := num/3
	i, cur, n := 0, 0, len(A)
	for i < n {
		cur += A[i]
		if cur == target {
			break
		}
		i++
	}
	if cur != target {
		return false
	}
	j := i + 1
	//最后一部分不能为空
	for j+1 < n {
		cur += A[j]
		if cur == target*2 {
			return true
		}
		j++
	}
	return false
}

//方法二
func CanThreePartsEqualSum1(A []int) bool {
	num := 0
	for _, v := range A {
		num += v
	}
	if num % 3 != 0 {
		return false
	}
	target := num/3
	cur := 0
	for i := 0; i < len(A); i++ {
		cur += A[i]
		if cur == target {
			for j := i+1; j < len(A); j++ {
				cur += A[j]
				if cur == target*2 {
					return true
				}
			}
			break
		}
	}
	return false
}