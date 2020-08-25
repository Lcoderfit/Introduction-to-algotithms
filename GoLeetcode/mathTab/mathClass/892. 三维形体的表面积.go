package mathClass


//一：O(n^2)
func SurfaceArea(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	dir := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	ret := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] != 0 {
				//该位置贡献上下两个表面积
				ret += 2
				//该位置与上下左右四个方向比较，只有大于某一个方向的立方体数量，才在该方向贡献表面积
				for _, d := range dir {
					iNew, jNew := i+d[0], j+d[1]
					tmp := 0
					if iNew >= 0 && iNew < rows && jNew >= 0 && jNew < cols {
						tmp = grid[iNew][jNew]
					}
					ret += max(grid[i][j]-tmp, 0)
				}
			}
		}
	}
	return ret
}

//二：O(n^2)
