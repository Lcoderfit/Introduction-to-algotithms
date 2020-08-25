package dp

import "math"

//一：贪心算法
func Jump(nums []int) int {
	length := len(nums)
	begin, end := 0, 1
	farthest, ans := 0, 0
	//其实位置为下标0的位置，所以初始化时begin, end的范围[0, 1)
	for end < length {
		for i := begin; i < end; i++ {
			farthest = max(farthest, i+nums[i])
		}
		//更新新的遍历范围为前一个范围的右边界到最远距离
		begin = end
		end = farthest + 1
		ans++
	}
	return ans
}

//二：dp
func Jump2(nums []int) int {
    length := len(nums)
    dp := make([]int, length)
    for i := 1; i < length; i++ {
        dp[i] = math.MaxInt32
    }
    for i := 0; i < length; i++ {
    	//j < length是为了防止i+nums[i]超出nums的长度范围
        for j := i+1; j < length && j <= i + nums[i]; j++ {
            dp[j] = min(dp[j], dp[i]+1)
        }
    }
    return dp[len(nums) - 1]
}

//三：固定for循环遍历
func Jump3(nums []int) int {
	//end记录当前“潜力子”所能到的最远位置
	farthest, end := 0, 0
	ret := 0
	for i := 0; i < len(nums) - 1; i++ {
		farthest = max(farthest, i+nums[i])
		if end == i {
			end = farthest
			//跳一步
			ret++
		}
	}
	return ret
}