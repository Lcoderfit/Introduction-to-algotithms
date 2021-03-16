package dp

//方法一：farthest表示当前所能到达的最大位置
func CanJump(nums []int) bool {
	farthest := 0
	for i := 0; i < len(nums)-1; i++ {
		farthest = max(farthest, i+nums[i])
		if farthest <= i {
			return false
		}
	}
	return farthest >= len(nums)-1
}

//方法二
func CanJump2(nums []int) bool {
	farthest := 0
	for i := 0; i < len(nums); i++ {
		if farthest < i {
			return false
		}
		if farthest < i+nums[i] {
			farthest = i + nums[i]
		}
	}
	return true
}
