package arr


// func removeDuplicates(nums []int) int {
//     slow, fast := 0, 1
//     for fast < len(nums) {
//         if nums[slow] != nums[fast] {
//             slow++
//             nums[slow] = nums[fast]
//         }
//         fast++
//     }
//     return slow+1
// }
func removeDuplicates(nums []int) int {
	length := len(nums)
	n := 1
	for i := 1; i < length; i++ {
		if nums[i] != nums[i-1] {
			nums[n] = nums[i]
			n++
		}
	}
	return n
}