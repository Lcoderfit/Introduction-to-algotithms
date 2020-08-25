package dp


func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Create2DSlice(rows, cols int) [][]int {
	arr := make([][]int, 0)
	for i := 0; i < rows; i++ {
		t := make([]int, cols)
		arr = append(arr, t)
	}
	return arr
}