package doublePtr


/*
双指针解法-端式双指针
时间复杂度：O(n)，n是nums的长度
空间复杂度：O(1)。
*/
func removeElement(nums []int, val int) int {
    i := 0
    n := len(nums)
    for i < n {
        if nums[i] == val {
            nums[i] = nums[n-1]
            n--
        } else {
            i++
        }
    }
    return n
}


/*
双指针解法-左式前后指针, 更好
时间复杂度：O(n)，n是nums的长度
空间复杂度：O(1)。
*/
func RemoveElement2(nums []int, val int) int {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[j] = nums[i]
			j++
		}
	}
	return j
}