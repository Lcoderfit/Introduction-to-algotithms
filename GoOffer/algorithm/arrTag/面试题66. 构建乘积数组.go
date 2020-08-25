package arrTag



//方法一：对称数组
func ConstructArr(a []int) []int {
	if len(a) <= 1 {
		return []int{}
	}
	ret := make([]int, len(a))
	ret[0] = 1
	for i := 1; i < len(a); i++ {
		//ret[i]表示i位置左边所有元素之乘，ret[i] = a[0]*a[1]...*a[i-1]
		ret[i] = ret[i-1]*a[i-1]
	}
	r := 1
	for j := len(a)-1; j >= 0; j-- {
		//ret[j]表示a中除去a[i]之后所有元素的乘积
		ret[j] = ret[j]*r
		r *= a[j]
	}
	return ret
}

//方法二：对称数组优化
func ConstructArr1(a []int) []int {
	if len(a) <= 1 {
		return []int{}
	}

	ret := make([]int, len(a))
	left := 1
	for i := range a {
		ret[i] = left
		left *= a[i]
	}
	right := 1
	for j := len(a) - 1; j >= 0; j-- {
		ret[j] *= right
		right *= a[j]
	}
	return ret
}

//方法三：双数组
func ConstructArr2(a []int) []int {
	if len(a) <= 1 {
		return []int{}
	}

	left := make([]int, len(a))
	right := make([]int, len(a))
	left[0] = 1
	for i := 1; i < len(a); i++ {
		left[i] = left[i-1]*a[i-1]
	}
	right[len(a)-1] = 1
	for j := len(a) - 2; j >= 0; j-- {
		right[j] = right[j+1]*a[j+1]
	}

	ret := make([]int, len(a))
	for i := 0; i < len(a); i++ {
		ret[i] = left[i]*right[i]
	}
	return ret
}