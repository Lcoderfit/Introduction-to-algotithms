package arr


func CheckPossibility(nums []int) bool {
	if len(nums) <= 2 {
		return true
	}
	//在第i个位置的元素，如果比i+1位置的大，那么要么将nums[i]改为nums[i-1]
	//要么将nums[i+1]改成nums[i]
	count := 0
	for i := 0; i < len(nums) - 1; i++ {
		if nums[i] > nums[i+1] {
			count++
			if count > 1 {
				return false
			}
			tmp := nums[i]
			//判断i的值，处理边界问题
			if i >= 1 {
				nums[i] = nums[i-1]
			} else {
				//当i为第一个元素时，直接将nums[i]变成更小的元素(nums[i] < nums[i+1])
				nums[i] = nums[i+1]
			}

			//修改nums[i]不可行，应该修改nums[i+1]
			if nums[i] > nums[i+1] {
				nums[i] = tmp
				nums[i+1] = nums[i]
			}
		}
	}
	return true
}