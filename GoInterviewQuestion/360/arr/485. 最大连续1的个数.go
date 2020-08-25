package arr


func findMaxConsecutiveOnes(nums []int) int {
	ret, count := 0, 0
	for _, v := range nums {
		if v == 1 {
			count++
			if ret < count {
				ret = count
			}
		} else {
			count = 0
		}
	}
	return ret
}