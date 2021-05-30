/*

方法1：
时间复杂度：O(n)
空间复杂度：O(1)

case1:

r:

*/
package medium

func MinCost1(costs [][]int) int {
	if costs == nil || len(costs) == 0 {
		return 0
	}
	var r, b, g = costs[0][0], costs[0][1], costs[0][2]
	for i := 1; i < len(costs); i++ {
		r, b, g = costs[i][0]+Min(b, g), costs[i][1]+Min(r, g), costs[i][2]+Min(r, b)
	}
	return Min(r, b, g)
}
