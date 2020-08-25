package bitwise


//位运算：记住num&(num-1)去除最低位的骚操作
func hammingWeight(num uint32) int {
	ret := 0
	for num != 0 {
		num = num&(num-1)
		ret++
	}
	return ret
}