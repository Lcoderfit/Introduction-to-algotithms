package remember


func DivingBoard(shorter int, longer int, k int) []int {
	if shorter == longer {
		return []int{k*shorter}
	}
	ret := make([]int, k + 1)
	for i := 0; i <= k; i++ {
		ret = append(ret, (k-i)*shorter + i*longer)
	}
	return ret
}
