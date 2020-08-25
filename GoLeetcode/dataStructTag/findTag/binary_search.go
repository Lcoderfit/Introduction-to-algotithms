package findTag


func BinarySearch(nums []int, target int) int {
	i, j := 0, len(nums) - 1
	for i <= j {
		mid := (i+j)/2
		if nums[mid] < target {
			i = mid + 1
		} else if nums[mid] > target {
			j = mid - 1
		} else {
			return mid
		}
	}
	return -1
}