package dp


func LastRemaining(n int, m int) int {
	//当只有一个人时，返回第一个人的编号，即0
	flag := 0
	for i := 2; i <= n; i++ {
		//k == m - 1 ==> m = k+1
		//y = (x+k+1)%n => y = (x+m)%n
		flag = (flag+m)%i
	}
	return flag
}