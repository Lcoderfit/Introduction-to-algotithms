package medium

import "sort"

func eraseOverlapIntervals(intervals [][]int) int {
	if intervals == nil || len(intervals) == 0 {
		return 0
	}

	// 通过左区间排序，然后根据前面的元素的右边界是否小于等于 后面元素的左边界来判断是否是可以构成符合要求的区间
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	n := len(intervals)
	// dp[i]表示以列表i结尾的最长的符合要求的区间长度
	dp := make([]int, n+1)
	// 注意初始化条件，区间长度至少为1，即当前的元素
	for i := range dp {
		dp[i] = 1
	}

	// 区间dp
	for i := 1; i <= n; i++ {
		for j := 1; j < i; j++ {
			if intervals[j-1][1] <= intervals[i-1][0] {
				dp[i] = Max(dp[i], dp[j]+1)
			}
		}
	}
	return n - dp[n]
}
