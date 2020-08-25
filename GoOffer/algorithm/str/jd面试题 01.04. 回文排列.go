package str


func CanPermutePalindrome(s string) bool {
	//利用异或运算,位向量
	var low, high int64
	for _, v := range s {
		//如果字母出现次数为偶数，则在对应位上最后异或结果为0
		//注意是>=, 低位0-63, 高位64-127
		if v >= 64 {
			high ^= 1 << (v - 64)
		} else {
			low ^= 1 << v
		}
	}
	//只有一个字母可能有奇数个，有奇数个的字母异或之后
	//对应位上为1，一定是2的幂
	if (high & (high - 1)) == 0 && (low & (low - 1)) == 0 {
		return true
	}
	return false
}
