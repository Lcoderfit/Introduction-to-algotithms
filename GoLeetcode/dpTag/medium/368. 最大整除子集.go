/*

方法1：动态规划
时间复杂度：O(n^2)
空间复杂度：O(n)

case1:
r:

case2:
r:

case3:
r:
*/
package medium

import (
	"sort"
)

func largestDivisibleSubset(nums []int) []int {
	if len(nums) == 0 {
		return nil
	}
	sort.Ints(nums)

	// dp[i]表示以nums[i]结尾的最长[0..i-1]中最长的整除子集长度
	dp := make([]int, len(nums))

	// 正整数
	maxSize, maxVal := 1, nums[0]
	for i := 0; i < len(nums); i++ {
		// 每个dp初始化为1，因为至少都有一个，就是nums[i]本身
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i]%nums[j] == 0 {
				dp[i] = Max(dp[i], dp[j]+1)
			}
		}
		if dp[i] > maxSize {
			maxSize, maxVal = dp[i], nums[i]
		}
	}

	// 先求出最长整除子集长度和整除子集最大值，再逆序遍历，求出整除子集的元素
	result := make([]int, 0)
	for i := len(nums) - 1; i >= 0 && maxSize > 0; i-- {
		if (dp[i] == maxSize) && (maxVal%nums[i] == 0) {
			maxSize--
			maxVal = nums[i]
			result = append(result, nums[i])
		}
	}
	return result
}

func largestDivisibleSubset1(nums []int) []int {
	if len(nums) == 0 {
		return nil
	}
	sort.Ints(nums)

	resultList := make([][]int, len(nums))
	for i := range nums {
		resultList[i] = []int{nums[i]}
	}

	for i := 0; i < len(nums); i++ {
		// 每个dp初始化为1，因为至少都有一个，就是nums[i]本身
		for j := 0; j < i; j++ {
			if nums[i]%nums[j] == 0 && len(resultList[j])+1 > len(resultList[i]) {
				// 如果是resultList[i] = append(resultList[j], []int{nums[i]})，由于append可能不会改变resultList[j]的地址，
				// 所以有可能会导致resultList[i]与resultList[j]变为同一个切片，即修改resultList[i]会影响到resultList[j]
				resultList[i] = append([]int{nums[i]}, resultList[j]...)
			}
		}
	}

	result := resultList[0]
	maxSize := 1
	for _, v := range resultList {
		if maxSize < len(v) {
			maxSize = len(v)
			result = v
		}
	}
	return result
}
