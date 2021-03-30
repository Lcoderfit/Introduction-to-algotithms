package medium

// Min 求最小值
func Min(nums ...int) int {
	// 32位整型最大小值
	result := nums[0]
	for _, num := range nums {
		if result > num {
			result = num
		}
	}
	return result
}

// Max 求最大值
func Max(nums ...int) int {
	result := nums[0]
	for _, num := range nums {
		if result < num {
			result = num
		}
	}
	return result
}
