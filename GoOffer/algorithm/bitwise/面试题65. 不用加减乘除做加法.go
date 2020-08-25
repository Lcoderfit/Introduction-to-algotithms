package bitwise


//^: 本位 &：进位
//最后结果为代表本位的变量，最后为0的为代表进位的变量
//进位循环一次需要左移动一位
func Add(a int, b int) int {
	for a != 0 {
		a, b = a&b, a^b
		a <<= 1
	}
	return b
}