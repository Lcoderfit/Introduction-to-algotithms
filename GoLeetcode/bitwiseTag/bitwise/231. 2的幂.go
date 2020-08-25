package bitwise


//n&(n-1)的活用，如果是2的幂次，说明二进制表示中只有一位，而n*(n-1)可以消除最低位
//如果只有一位的话，消除之后就为0
func IsPowerOfTwo(n int) bool {
	return n > 0 && (n&(n-1)) == 0
}