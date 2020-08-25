package bitwise


//不用加减乘除法、for、while、if、else、switch、case等关键字及条件判断语句（A?B:C）。
func sumNums(n int) int {
	//短路原理，如果n==0, 则不会再递归调用plus()
	//n > 0,则会调用plus
	_ = n > 0 && add(&n, sumNums(n - 1))
	return n
}

func add(n *int, b int) bool {
	*n += b
	return true
}



//不用加减乘除
//方法一：位运算
func SumNums1(n int) int {
	ret := 0
	for i := 1; i <= n; i++ {
		ret = add1(ret, i)
	}
	return ret
}

func add1(a, b int) int {
	for b != 0 {
		a, b = a^b, a&b
		b <<= 1
	}
	return a
}