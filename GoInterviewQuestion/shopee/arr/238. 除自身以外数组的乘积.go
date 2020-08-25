package arr


func productExceptSelf(nums []int) []int {
	//对称数组
	ret := make([]int, len(nums))
	//热更新初始化法
	left := 1
	for i := 0; i < len(nums); i++ {
		ret[i] = left
		left *= nums[i]
	}

	right := 1
	for j := len(nums) - 1; j >= 0; j-- {
		ret[j] *= right
		right *= nums[j]
	}
	return ret
}