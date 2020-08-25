package arrTag


//双指针
func TwoSum(nums []int, target int) []int {
	i, j := 0, len(nums) - 1
	ret := make([]int, 0)
	for i < j {
		if nums[i] + nums[j] == target {
			ret = append(ret, nums[i], nums[j])
			break
		} else if nums[i] + nums[j] > target {
			j--
		} else {
			i++
		}
	}
	return ret
}