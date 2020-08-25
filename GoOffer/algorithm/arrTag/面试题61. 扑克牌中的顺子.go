package arrTag

import "sort"

func IsStraight(nums []int) bool {
	//对slice进行排序
	sort.Ints(nums)
	sub := 0
	for i := 0; i < 4; i++ {
		if nums[i] == 0 {
			continue
		}
		if nums[i] == nums[i+1] {
			return false
		}
		sub += nums[i+1] - nums[i]
	}
	return sub < 5
}