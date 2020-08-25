package arr


// func findDuplicate(nums []int) int {
//     if len(nums) <= 1 {
//         return -1
//     }
//     sort.Ints(nums)
//     for i := 1; i < len(nums); i++ {
//         if nums[i] == nums[i-1] {
//             return nums[i]
//         }
//     }
//     return -1
// }


func findDuplicate(nums []int) int {
	//转化为快慢指针寻找链表中有环，值相当于地址
	//快指针路程需是慢指针的两倍，所以初始化时fast在slow后面
	slow, fast := nums[0], nums[nums[0]]
	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}

	ret := 0
	for ret != slow {
		ret = nums[ret]
		slow = nums[slow]
	}
	return ret
}