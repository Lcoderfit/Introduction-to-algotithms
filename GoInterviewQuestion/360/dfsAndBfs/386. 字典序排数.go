package dfsAndBfs


func lexicalOrder(n int) []int {
	ret := make([]int, 0)
	//构建十叉树
	//cur表示当前的节点
	var dfs func(cur, n int)
	dfs = func(cur, n int) {
		if cur > n {
			return
		}
		ret = append(ret, cur)
		for i := 0; i < 10; i++ {
			if cur*10 + i > n {
				return
			}
			dfs(cur*10+i, n)
		}
	}
	for i := 1; i < 10; i++ {
		dfs(i, n)
	}
	return ret
}