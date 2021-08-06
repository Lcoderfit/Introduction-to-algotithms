package dp

import "fmt"

/*
时间：O(n^2)
空间：O(n)
方法一：dp


方法二：dp+二分
时间：O(nlogn)
空间：O(n)
*/
func LengthOfLIS1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	// dp[i]表示长度为i+1的末尾值最小的LIS
	d := make([]int, 0)
	for _, v := range nums {
		if len(d) == 0 || v > d[len(d)-1] {
			d = append(d, v)
		} else {
			l, r := 0, len(d)-1
			for l < r {
				mid := (l + r) / 2
				if v > d[mid] {
					l = mid + 1
				} else {
					r = mid
				}
			}
			d[l] = v
		}
	}
	return len(d)
}

func LengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			} else {
				dp[i] = max(dp[i], dp[j])
			}
		}
	}
	fmt.Println(dp)
	return dp[len(nums)-1]
}
