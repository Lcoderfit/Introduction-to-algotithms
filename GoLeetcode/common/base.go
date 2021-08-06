/*

方法1：
时间复杂度：O()
空间复杂度：O()

case1:

r:

*/
package common

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

// 求总和
func Sum(nums ...int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	ans := 0
	for _, v := range nums {
		ans += v
	}
	return ans
}
