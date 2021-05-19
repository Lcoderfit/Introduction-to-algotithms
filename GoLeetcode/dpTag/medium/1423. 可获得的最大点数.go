/*

方法1：滑动窗口(从右往左滑动)
时间复杂度：O(n)
空间复杂度：O(1)

方法2和3：滑动窗口（从左往右滑动）
时间复杂度：O(n)
空间复杂度：O(1)

方法4：滑动窗口+前缀和(空间复杂度比较高，暂不写)
时间复杂度：O(n)
空间复杂度：O(n)

case1:

r:

*/
package medium

import "math"

func MaxScore(cardPoints []int, k int) int {
	if cardPoints == nil || len(cardPoints) == 0 || k == 0 {
		return 0
	}
	sum := 0
	for i := 0; i < k; i++ {
		sum += cardPoints[i]
	}
	ans := sum
	n := len(cardPoints)
	for i := 0; i < k; i++ {
		// 滑动窗口，从右往左滑动，滑动时先减去右窗口右边的元素，然后加上左窗口左边的元素
		sum += cardPoints[n-i-1] - cardPoints[k-i-1]
		ans = Max(ans, sum)
	}
	return ans
}

func MaxSorce2(cardPoints []int, k int) int {
	if cardPoints == nil || len(cardPoints) == 0 || k == 0 {
		return 0
	}
	n := len(cardPoints)
	windowsSize := n - k
	sum := 0
	windowsSum := 0
	ans := math.MaxInt32
	for i := 0; i < n; i++ {
		sum += cardPoints[i]
		windowsSum += cardPoints[i]
		if i > windowsSize-1 {
			windowsSum -= cardPoints[i-windowsSize]
			ans = Min(ans, windowsSum)
		}
	}
	return sum - ans
}

func MaxScore3(cardPoints []int, k int) int {
	if cardPoints == nil || len(cardPoints) == 0 || k == 0 {
		return 0
	}

	// 计算元素总和
	sum := Sum(cardPoints...)
	// 计算初始窗口元素和
	windowsSize := len(cardPoints) - k
	windowsSum := Sum(cardPoints[:windowsSize]...)
	minWindowsSum := 0

	// 窗口进行滑动，取最小窗口元素和，用所有元素总和减去最小窗口和即为符合题意的最大值
	minWindowsSum = windowsSum
	for i := windowsSize; i < len(cardPoints); i++ {
		windowsSum += cardPoints[i] - cardPoints[i-windowsSize]
		minWindowsSum = Min(minWindowsSum, windowsSum)
	}
	return sum - minWindowsSum
}
