/*
方法1：动态规划
时间复杂度：O(n)
空间复杂度：O(n)

方法2：动态规划+滚动数组
时间复杂度：O(n)
空间复杂度：O(1)

方法3：贪心算法:该问题的关键在于选择波峰与波谷，计算波峰与波谷的个数即可
时间复杂度：O(n)
空间复杂度：O(1)


case1:
1 7 4 9 2 5
r:
6

case2:
0 0
r:
1


*/
package medium

// WiggleMaxLength 摆动子序列的最大长度
func WiggleMaxLength(nums []int) int {
	if nums == nil {
		return 0
	}
	n := len(nums)
	up := make([]int, n)
	down := make([]int, n)
	up[0], down[0] = 1, 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			up[i] = Max(up[i-1], down[i-1]+1)
			down[i] = down[i-1]
		} else if nums[i] == nums[i-1] {
			up[i] = up[i-1]
			down[i] = down[i-1]
		} else {
			up[i] = up[i-1]
			down[i] = Max(down[i-1], up[i-1]+1)
		}
	}
	return Max(down[n-1], up[n-1])
}

func WiggleMaxLength2(nums []int) int {
	if nums == nil {
		return 0
	}
	up, down := 1, 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			up = Max(up, down+1)
		} else if nums[i] < nums[i-1] {
			down = Max(down, up+1)
		}
	}
	return Max(up, down)
}

// WiggleMaxLength3 该问题的关键在于选择波峰与波谷，计算波峰与波谷的个数即可
func WiggleMaxLength3(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return 1
	}
	ans := 1
	preDiff := nums[1] - nums[0]
	if preDiff != 0 {
		ans++
	}
	for i := 2; i < len(nums); i++ {
		diff := nums[i] - nums[i-1]
		if diff > 0 && preDiff <= 0 || diff < 0 && preDiff >= 0 {
			ans++
			preDiff = diff
		}
	}
	return ans
}
