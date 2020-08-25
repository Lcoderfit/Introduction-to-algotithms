package arr

import (
	"math"
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return nil
	}
	sort.Ints(candidates)
	return dfs(candidates, target)
}

func dfs(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return [][]int{}
	}
	ret := make([][]int, 0)
	first := math.MinInt32
	for i := 0; i < len(candidates); i++ {
		//用来过滤重复的数据
		if first == candidates[i] {
			continue
		}
		first = candidates[i]
		if candidates[i] > target {
			break
		} else if candidates[i] < target {
			//切片左右边界最大允许值x <= len(arr)
			tailArr := combinationSum2(candidates[i+1:], target - candidates[i])
			for _, v := range tailArr {
				//不要求输出数据有序，可以直接append压入
				ret = append(ret, append(v, candidates[i]))
			}
		} else {
			ret = append(ret, []int{candidates[i]})
			break
		}
	}
	return ret
}