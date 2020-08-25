package bitwise

func SingleNumbers(nums []int) []int {
	t := 0
	for i := 0; i < len(nums); i++ {
		t ^= nums[i]
	}
	p := 1
	for p & t == 0 {
		p <<= 1
	}

	left, right := 0, 0
	for _, v := range nums {
		if v & p > 0 {
			left ^= v
		} else {
			right ^= v
		}
	}
	return []int{left, right}
}