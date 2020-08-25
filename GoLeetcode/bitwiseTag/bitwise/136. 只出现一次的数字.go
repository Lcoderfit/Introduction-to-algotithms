package bitwise


/*
时间复杂度：O(n)，其中 n是数组长度。只需要对数组遍历一次。

空间复杂度：O(1)。
 */
func singleNumber(nums []int) int {
	ret := 0
	for _, v := range nums {
		ret ^= v
	}
	return ret
}