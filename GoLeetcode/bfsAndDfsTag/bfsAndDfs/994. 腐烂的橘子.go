/*
时间复杂度：O()
空间复杂度：O()
 */
package bfsAndDfs

func OrangesRotting(grid [][]int) int {
	if len(grid) < 1 || len(grid[0]) < 1 {
		return -1
	}

	rows, cols := len(grid), len(grid[0])
	//保存最小分钟数
	minutes := 0
	directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	queue := make([][]int, 0)

	//将所有烂橘子所在的位置放入queue中
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 2 {
				queue = append(queue, []int{i, j, 0})
			}
		}
	}

	//多源BFS
	for len(queue) != 0 {
		//从queue中弹出烂橘子所在的位置
		q := queue[0]
		i, j := q[0], q[1]
		minutes = q[2]
		//弹出节点
		queue = queue[1:]
		for _, d := range directions {
			iNew, jNew := i+d[0], j+d[1]
			if iNew >= 0 && iNew < rows && jNew >= 0 && jNew < cols && grid[iNew][jNew] == 1 {
				grid[iNew][jNew] = 2
				queue = append(queue, []int{iNew, jNew, minutes + 1})
			}
		}
	}

	//经过BFS之后，判断grid中是否还含有新鲜橘子
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 1 {
				return -1
			}
		}
	}

	return minutes
}
