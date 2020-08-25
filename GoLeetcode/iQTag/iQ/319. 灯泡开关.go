/*
时间复杂度：O(1)
空间复杂度：O(1)
*/
package iQ

import "math"

func BulbSwitch(n int) int {
	if n <= 0 {
		return 0
	}
	t := math.Sqrt(float64(n))
	return int(t)
}
