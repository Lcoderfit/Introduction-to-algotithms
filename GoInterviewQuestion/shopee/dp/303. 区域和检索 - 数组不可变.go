package dp


type NumArray struct {
	dp []int
}


func Constructor(nums []int) NumArray {
	length := len(nums)
	dp := make([]int, length + 1)
	//dp[i]为nums[0...i-1]的前缀和
	for i := 1; i <= length; i++ {
		dp[i] = dp[i-1] + nums[i-1]
	}
	return NumArray{dp: dp}
}


func (this *NumArray) SumRange(i int, j int) int {
	//sumRange(i, j)是第i+1个元素到第j+1个元素之和
	//也就等于前j+1个元素的前缀和减去前i个元素的前缀和
	return this.dp[j+1] - this.dp[i]
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */