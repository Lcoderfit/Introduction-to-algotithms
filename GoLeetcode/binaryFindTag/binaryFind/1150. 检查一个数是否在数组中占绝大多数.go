package binaryFind

/*
方法1：
时间复杂度：O(logn)
空间复杂度：O(1)

case1:
9 5
2 4 5 5 5 5 5 6 6

case2:
4 101
10 100 101 101
*/

func IsMajorityElement(nums []int, target int) bool {
	left := binarySearchLeft(nums, target)
	return (left != -1) && (left + len(nums) / 2 < len(nums)) && (nums[left + len(nums) / 2] == target)
}

func binarySearchLeft(nums []int, target int) int {
	i, j := 0, len(nums) - 1
	for i <= j {
		mid := (i + j) / 2
		if nums[mid] >= target {
			j = mid - 1
		} else {
			i = mid + 1
		}
	}
	if (i < len(nums)) && (nums[i] == target) {
		return i
	}
	return -1
}