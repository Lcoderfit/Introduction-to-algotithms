/*

方法1：替换法
时间复杂度：O()
空间复杂度：O()

方法2：滑动窗口
时间复杂度：O()
空间复杂度：O()

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
		// 最右边的一个元素，需要用左边的第k-i-1位置的元素来换，
		// 因为如果右边只取了一个，则左边肯定取了k-1个，由于是从外取到里面，所以肯定是第k个取不到
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
		if i > windowsSize - 1 {
			windowsSum -= cardPoints[i-windowsSize]
			ans = Min(ans, windowsSum)
		}
	}
	return sum - ans
}
