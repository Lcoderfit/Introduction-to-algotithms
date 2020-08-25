package arr


//方法一：高斯定理
func MissingNumber(nums []int) int {
	length := len(nums)
	sum := length*(length+1)/2
	for _, v := range nums {
		sum -= v
	}
	return sum
}

//二：位运算
func MissingNumber2(nums []int) int {
	ret := 0
	//异或运算，相同则抵消为0，下标[1-n], 数字为[1-n]缺少一个元素
	for i, v := range nums {
		ret ^= (i+1)^v
	}
	return ret
}