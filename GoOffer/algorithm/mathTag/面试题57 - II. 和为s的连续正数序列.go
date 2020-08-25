/*
时间复杂度：O(target)
空间复杂度：O(1)
 */
package mathTag

func FindContinuousSequence(target int) [][]int {
	if target == 1 {
		return [][]int{{1}}
	}
	ret := make([][]int, 0)
	i, j := 1, 2
	sum := i + j
	for i < target/2 {
		if sum == target {
			ret = append(ret, genContinuousSequence(i, j))
			sum -= i
			i++
		} else if sum < target {
			j++
			sum += j
		} else {
			sum -= i
			i++
		}
	}
	return ret
}

func genContinuousSequence(i, j int) []int {
	ret := make([]int, 0)
	for k := i; k <= j; k++ {
		ret = append(ret, k)
	}
	return ret
}
