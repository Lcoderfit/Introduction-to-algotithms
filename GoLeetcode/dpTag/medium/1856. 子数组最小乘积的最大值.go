/*

方法1：双单调栈+前缀和
时间复杂度：O(n)
空间复杂度：O(n)

case1:

r:

*/
package medium

import "math"

// 子数组最小乘积的最大值
func MaxSumMinProduct(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	nums = append(append([]int{0}, nums...), 0)
	// 前缀和数组
	preSum := []int{0}
	for _, v := range nums {
		preSum = append(preSum, v+preSum[len(preSum)-1])
	}

	// 单调栈，找到i右边最近的一个小于nums[i]的数的索引
	stack := make([]int, 0)
	rightFirstSmaller := make([]int, len(nums))
	for i := range nums {
		for len(stack) > 0 && nums[i] < nums[stack[len(stack)-1]] {
			rightFirstSmaller[stack[len(stack)-1]] = i
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	// 找到i左边最近的一个小于nums[i]的数的索引
	stack = make([]int, 0)
	leftFirstSmaller := make([]int, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		for len(stack) > 0 && nums[i] < nums[stack[len(stack)-1]] {
			leftFirstSmaller[stack[len(stack)-1]] = i
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	res := 0
	for i := 1; i < len(nums)-1; i++ {
		left := leftFirstSmaller[i]
		right := rightFirstSmaller[i]
		res = Max(res, nums[i]*(preSum[right]-preSum[left+1]))
	}
	return res % (int(math.Pow(10, 9) + 7))
}
