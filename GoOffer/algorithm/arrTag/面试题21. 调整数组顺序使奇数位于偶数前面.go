package arrTag


func Exchange(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	i, j := 0, 0
	for j < len(nums) {
		if nums[j] % 2 == 1 {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
		j++
	}
	return nums
}