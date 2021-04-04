/*
方法1：动态规划
时间复杂度：O(n)
空间复杂度：O(n)

case1:

r:

*/
package medium

func countBits(num int) []int {
	ans := []int{0}
	for i := 1; i <= num; i++ {
		count := 0
		t := i
		for t > 0 {
			count++
			t = (t - 1) & t
		}
		ans = append(ans, count)
	}
	return ans
}
