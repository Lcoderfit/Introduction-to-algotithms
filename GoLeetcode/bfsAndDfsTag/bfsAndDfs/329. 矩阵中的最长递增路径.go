package bfsAndDfs

var dir = [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

//方法一：拓扑排序，
/*
时间复杂度 : O(mn)。拓扑排序的时间复杂度为 O(V+E) = O(mn)O(V+E)=O(mn)。
VV 是顶点总数，EE 是边总数。本问题中，O(V) = O(mn)O(V)=O(mn)，O(E) = O(4V) = O(mn)O(E)=O(4V)=O(mn)。
空间复杂度 : O(mn)O(mn)。我们需要存储出度和每层的叶子。
 */
func LongestIncreasingPath(matrix [][]int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	//计算所有格点的出度
	rows, cols := len(matrix), len(matrix[0])
	outDegree := createSlice(rows, cols)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			for _, d := range dir {
				if check(i+d[0], j+d[1], rows, cols) {
					if matrix[i][j] < matrix[i+d[0]][j+d[1]] {
						outDegree[i][j]++
					}
				}
			}
		}
	}

	//找出出度为0的格点，存入slice中
	leaves := make([][]int, 0)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if outDegree[i][j] == 0 {
				leaves = append(leaves, []int{i, j})
			}
		}
	}

	//反向拓扑，累计最大的路径长度
	height := 0
	for len(leaves) != 0 {
		height++
		newLeaves := make([][]int, 0)
		for _, node := range leaves {
			x, y := node[0], node[1]
			for _, d := range dir {
				iNew, jNew := x+d[0], y+d[1]
				if check(iNew, jNew, rows, cols) && matrix[iNew][jNew] < matrix[x][y] {
					outDegree[iNew][jNew]--
					if outDegree[iNew][jNew] == 0 {
						newLeaves = append(newLeaves, []int{iNew, jNew})
					}
				}
			}
		}
		leaves = newLeaves
	}
	return height
}

func createSlice(rows, cols int) [][]int {
	ret := make([][]int, 0)
	for i := 0; i < rows; i++ {
		t := make([]int, cols)
		ret = append(ret, t)
	}
	return ret
}

func check(i, j int, rows, cols int) bool {
	if i >= 0 && i < rows && j >= 0 && j < cols {
		return true
	}
	return false
}


//方法二：记忆化
func LongestIncreasingPath1(matrix [][]int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	rows, cols := len(matrix), len(matrix[0])
	cache := createSlice(rows, cols)
	ret := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			ret = max(ret, dfs(matrix, i, j, cache))
		}
	}
	return ret
}

func dfs(matrix [][]int, i, j int, cache [][]int) int {
	if cache[i][j] != 0 {
		return cache[i][j]
	}
	rows, cols := len(matrix), len(matrix[0])
	for _, d := range dir {
		iNew, jNew := i + d[0], j + d[1]
		if iNew >= 0 && iNew < rows && jNew >= 0 && jNew < cols {
			if matrix[i][j] < matrix[iNew][jNew] {
				//保存上下左右四个方向的最长路径值(不包括matrix[i][j]本身)
				cache[i][j] = max(cache[i][j], dfs(matrix, iNew, jNew, cache))
			}
		}
	}
	cache[i][j]++
	return cache[i][j]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}