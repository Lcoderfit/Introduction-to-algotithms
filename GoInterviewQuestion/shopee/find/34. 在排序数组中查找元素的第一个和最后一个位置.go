package find


func SearchRange(nums []int, target int) []int {
	length := len(nums)
	if length == 0 || target > nums[length - 1] || target < nums[0] {
		return []int{-1, -1}
	}

	ret := make([]int, 0)
	i, j := 0, length - 1
	//找到最左边出现的位置
	for i < j {
		mid := (i+j)/2
		if target <= nums[mid] {
			j = mid
		} else {
			i = mid + 1
		}
	}
	if nums[i] != target {
		return []int{-1, -1}
	}
	ret = append(ret, i)

	i, j = 0, len(nums) - 1
	//找到最右边出现的位置
	for i < j {
		mid := (i+j+1)/2
		if target >= nums[mid] {
			i = mid
		} else {
			j = mid - 1
		}
	}
	ret = append(ret, i)
	return ret
}