package mathClass


func IsHappy(n int) bool {
	//计算一个数各个位数上的数字平方和
	countSquare := func(n int) int {
		sum := 0
		for n != 0 {
			sum += (n%10)*(n%10)
			n /= 10
		}
		return sum
	}

	//如果是快乐数，则组后会得到1，如果不是快乐数，则得到的结果会循环
	//快慢指针
	fast, slow := countSquare(n), n
	for fast != slow {
		fast = countSquare(fast)
		fast = countSquare(fast)
		slow = countSquare(slow)
	}
	return fast == 1
}