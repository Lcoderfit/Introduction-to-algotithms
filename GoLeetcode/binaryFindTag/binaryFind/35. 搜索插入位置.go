package binaryFind
/*
35. 搜索插入位置.py
时间复杂度：O(logn)
空间复杂度：O(1)
*/

func SearchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	i, j := 0, len(nums) - 1
	for i <= j {
		mid := (i + j) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			j = mid - 1
		} else {
			i = mid + 1
		}
	}
	return i
}