package dp


func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func create2DArray(rows, cols int) [][]int {
	ret := make([][]int, 0)
	for i := 0; i < rows; i++ {
		t := make([]int, cols)
		ret = append(ret, t)
	}
	return ret
}