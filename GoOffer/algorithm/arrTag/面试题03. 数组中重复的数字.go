package arrTag


//一：hashmap
func FindRepeatNumber(nums []int) int {
	hash := make(map[int]int, len(nums))
	for _, v := range nums {
		if _, ok := hash[v]; !ok {
			hash[v]++
		} else {
			return v
		}
	}
	return -1
}

//二：slice代替map
func FindRepeatNumber2(nums []int) int {
	//如果key可以跟slice的小标索引对应，那么就可以用slice代替map
	hash := make([]int, len(nums))
	for _, v := range nums {
		hash[v]++
		if hash[v] > 1 {
			return v
		}
	}
	return 0
}

//三：原地调换位置
func FindRepeatNumber3(nums []int) int {
	i := 0
	for i < len(nums) {
		if nums[i] == i {
			i++
			//如果其他数据都与其下标对应，此时如果找到一个不与下标对应的数
			//nums[i] == x   nums[x] == x == nums[i]
		} else if nums[i] == nums[nums[i]] {
			return nums[i]
		} else {
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		}
	}
	return -1
}