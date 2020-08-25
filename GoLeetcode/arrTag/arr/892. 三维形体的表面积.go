package arr


func surfaceArea(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	ret := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] != 0 {
				//先将该位置的包括之后会重叠的所有表面积加上
				ret += 4*grid[i][j] + 2
			}
			//减去该位置与其左边重叠的部分的面积
			//重叠部分面积为两个位置高度的最小值*2，因为重叠的是两块面积，自己的和左边位置的
			if j > 0 {
				ret -= min(grid[i][j], grid[i][j-1])*2
			}
			//减去该位置与其后面重叠的部分的面积
			if i > 0 {
				ret -=min(grid[i][j], grid[i-1][j])*2
			}
		}
	}
	return ret
}