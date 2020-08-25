package arrTag

func SpiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}

	start := 0
	rows, cols := len(matrix), len(matrix[0])
	ret := make([]int, 0)
	for start*2 < rows && start*2 < cols {
		endx, endy := rows-start-1, cols-start-1
		for i := start; i <= endy; i++ {
			ret = append(ret, matrix[start][i])
		}
		//判断是否有右路
		if start < endx {
			for i := start + 1; i <= endx; i++ {
				ret = append(ret, matrix[i][endy])
			}
		}
		//判断是否有下路
		if start < endx && start < endy {
			for i := endy - 1; i >= start; i-- {
				ret = append(ret, matrix[endx][i])
			}
		}
		//判断是否有左路
		if start < endx-1 && start < endy {
			for i := endx - 1; i > start; i-- {
				ret = append(ret, matrix[i][start])
			}
		}

		start++
	}
	return ret
}
