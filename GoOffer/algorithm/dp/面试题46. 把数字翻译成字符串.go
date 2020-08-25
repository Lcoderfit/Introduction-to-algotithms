package dp


//方法一：递归解法
func TranslateNum(num int) int {
	//递归解法
	if num <= 9 {
		return 1
	}
	//截取后面两位
	ba := num%100
	//06或者28
	if ba <= 9 || ba >= 26 {
		return TranslateNum(num/10)
	}
	return TranslateNum(num/10) + TranslateNum(num/100)
}

//方法二：dp
func TranslateNum1(num int) int {
	dp := []int{1}
	f, b := 0, 0
	i := 0
	sum := 0
	for num != 0 {
		i++
		dp = append(dp, 0)
		f = num % 10
		dp[i] = dp[i-1]
		sum = f*10 + b
		if i > 1 && sum > 9 && sum < 26 {
			dp[i] += dp[i-2]
		}
		num /= 10
		b = f
	}
	return dp[i]
}

//方法三
