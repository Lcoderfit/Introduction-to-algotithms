package find


//二分查找，确定左右边界
func Search(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	l, r := 0, len(nums) - 1
	for l < r {
		mid := (l+r)/2
		if nums[mid] > target {
			r = mid - 1
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid
		}
	}
	if nums[l] != target {
		return 0
	}

	lIndex := l
	l, r = 0, len(nums) - 1
	for l < r  {
		mid := (l+r+1)/2
		if nums[mid] > target {
			r = mid - 1
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			l = mid
		}
	}
	return r - lIndex + 1
}