/*
时间复杂度：O(1)
空间复杂度：O(1)
*/
package iQ

import "sort"

func NumMovesStones(a int, b int, c int) []int {
	arr := []int{a, b, c}
	sort.Ints(arr)
	max := arr[2] - arr[0] - 2
	if arr[2] - arr[1] == 1 && arr[1] - arr[0] == 1 {
		return []int{0, max}
	}
	if arr[2] - arr[1] <= 2 || arr[1] - arr[0] <= 2 {
		return []int{1, max}
	}
	return []int{2, max}
}
