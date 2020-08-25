/*
1 <= n <= 10^5
时间复杂度：O(1)
空间复杂度：O(1)

1/n + (n-2)*(1/n)
*/
package iQ


func NthPersonGetsNthSeat(n int) float64 {
	if n == 1 {
		return 1
	}
	return 0.5
}
