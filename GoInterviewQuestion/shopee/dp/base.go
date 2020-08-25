package dp


func Create2DMatrix(m, n int) [][]int {
	ret := make([][]int, 0)
	for i := 0; i < m; i++ {
		t := make([]int, n)
		ret = append(ret, t)
	}
	return ret
}