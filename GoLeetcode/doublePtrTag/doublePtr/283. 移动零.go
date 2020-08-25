package doublePtr


/*
双指针解法1：一个指针用于for循环迭代，遍历所有元素，另一个指针i用于指向最开始的元素
当nums[j] != 0; 则i++

时间复杂度：O(n)，n是nums的长度
空间复杂度：O(1)。
*/
func MoveZeroes(nums []int)  {
	if len(nums) == 0 {
		return
	}
	//从第一个元素开始
	i, j := 0, 0
	for j < len(nums) {
		if nums[j] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
		j++
	}
}

/*
双指针解法2: 固定外层指针用于循环，另一个指针仅当交换发生时向后移动一个位置
时间复杂度：O(n)，n是nums的长度
空间复杂度：O(1)。
*/
func MoveZeroes1(nums []int)  {
	if len(nums) == 0 {
		return
	}
	i, j := 0, 0
	for j < len(nums) {
		if nums[j] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
		j++
	}
}