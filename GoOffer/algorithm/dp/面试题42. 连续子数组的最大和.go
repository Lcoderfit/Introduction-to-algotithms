package dp

import "math"

//方法一
func MaxSubArray(nums []int) int {
	sum := 0
	ret := math.MinInt32
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if ret < sum {
			ret = sum
		}
		if sum < 0 {
			sum = 0
		}
	}
	return ret
}

//方法二: dp
func MaxSubArray1(nums []int) int {
	ret := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > 0 {
			nums[i] = nums[i] + nums[i-1]
		}
		if ret < nums[i] {
			ret = nums[i]
		}
	}
	return ret
}