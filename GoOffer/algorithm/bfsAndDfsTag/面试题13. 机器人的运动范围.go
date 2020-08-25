package bfsAndDfsTag

import (
	"fmt"
)

func MovingCount(m int, n int, k int) int {
	queue, visited := [][]int{{0, 0, 0, 0}}, make([][]bool, 0)
	for i := 0; i < m; i++ {
		t := make([]bool, n)
		visited = append(visited, t)
	}
	ret := 0
	for len(queue) != 0 {
		node := queue[0]
		i, j, si, sj := node[0], node[1], node[2], node[3]
		queue = queue[1:]
		if i < 0 || i >= m || j < 0 || j >= n || visited[i][j] || (si+sj) > k {
			if (si + sj) > k {
				fmt.Println(i, j, si, sj)
			}
			continue
		}
		visited[i][j] = true
		ret++
		siNew, sjNew := si, sj
		if (i+1)%10 == 0 {
			siNew -= 8
		} else {
			siNew += 1
		}
		if (j+1)%10 == 0 {
			sjNew -= 8
		} else {
			sjNew += 1
		}
		queue = append(queue, []int{i + 1, j, siNew, sj}, []int{i, j + 1, si, sjNew})
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%v ", visited[i][j])
		}
		fmt.Println()
	}

	return ret
}

//方法二：内部定义函数变量的DFS
func movingCount1(m int, n int, k int) int {
	visited := make([][]bool, 0)
	for i := 0; i < m; i++ {
		t := make([]bool, n)
		visited = append(visited, t)
	}

	ret := 0
	//必须先声明dfs函数变量
	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || i >= m || j < 0 || j >= n || visited[i][j] || addNumb(i, j) > k {
			return
		}
		ret++
		visited[i][j] = true

		dfs(i+1, j)
		dfs(i, j+1)
	}
	dfs(0, 0)
	return ret
}

func addNumb1(a, b int) int {
	ret := 0
	for a != 0 {
		ret += a % 10
		a /= 10
	}
	for b != 0 {
		ret += b % 10
		b /= 10
	}
	return ret
}

//方法三：DFS
func MovingCount2(m int, n int, k int) int {
	visited := make([][]bool, 0)
	for i := 0; i < m; i++ {
		t := make([]bool, n)
		visited = append(visited, t)
	}
	ret := 0
	dfs2(visited, 0, 0, k, &ret)
	return ret
}

func dfs2(visited [][]bool, i, j, k int, ret *int) {
	if i < 0 || i >= len(visited) || j < 0 || j >= len(visited[0]) || visited[i][j] || addNumb(i, j) > k {
		return
	}
	*ret++
	visited[i][j] = true
	dfs2(visited, i+1, j, k, ret)
	dfs2(visited, i, j+1, k, ret)
}

func addNumb(a, b int) int {
	ret := 0
	for a != 0 {
		ret += a % 10
		a /= 10
	}
	for b != 0 {
		ret += b % 10
		b /= 10
	}
	return ret
}

//方法四：递归+回溯
func MovingCount3(m int, n int, k int) int {

	// 记录已走过的格子
	moved := make([][]bool, m)
	for i := 0; i < m; i++ {
		moved[i] = make([]bool, n)
	}

	nums := 0 //记录走过的全部格子总数

	// 递归 + 回溯
	var walk func(x, y int)
	walk = func(x, y int) {
		if countXY(x, y) > k { //判断此位置x，y 各位之和是否大于k
			return
		}
		moved[x][y] = true //标记为已走
		nums++             //走过的格子总数+1

		// 判断四周的情况
		if x-1 >= 0 && moved[x-1][y] == false { //向上走一步
			walk(x-1, y)
		}
		if y-1 >= 0 && moved[x][y-1] == false { //向左走一步
			walk(x, y-1)
		}
		if x+1 < m && moved[x+1][y] == false { //向下走一步
			walk(x+1, y)
		}
		if y+1 < n && moved[x][y+1] == false { //向右走一步
			walk(x, y+1)
		}

		//如果上面四周都不能走，则回退，回退没有任何操作
		//也没终止条件约束，最终计算出走过的格子总数即可
	}

	walk(0, 0) //从0，0出发递归
	return nums
}

// 计算x，y的个位数之和
func countXY(x, y int) (result int) {
	for x != 0 {
		result = result + (x % 10)
		x = x / 10
	}
	for y != 0 {
		result = result + (y % 10)
		y = y / 10
	}
	return
}
