package stackandqueue


func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return nil
	}

	index := make([]int, 0)
	for i := 0; i < k; i++ {
		for len(index) > 0 && nums[i] >= nums[index[len(index)-1]] {
			index = index[:len(index)-1]
		}
		index = append(index, i)
	}

	ret := make([]int, 0)
	for i := k; i < len(nums); i++ {
		ret = append(ret, nums[index[0]])

		for len(index) > 0 && nums[i] >= nums[index[len(index)-1]] {
			index = index[:len(index)-1]
		}
		if len(index) > 0 && index[0] <= (i-k) {
			index = index[1:]
		}
		index = append(index, i)
	}
	ret = append(ret, nums[index[0]])

	return ret
}
