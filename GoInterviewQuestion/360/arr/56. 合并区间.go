package arr

import "sort"

func merge(intervals [][]int) [][]int {
	//先将intervals中的sclie按照第一个元素的大小排序
	if len(intervals) <= 0 {
		return nil
	}
	sort.Slice(intervals, func(i, j int) bool {
		//从小到大排序
		return intervals[i][0] < intervals[j][0]
	})
	ret := make([][]int, 0)
	for _, v := range intervals {
		if len(ret) == 0 || ret[len(ret) - 1][1] < v[0] {
			ret = append(ret, v)
		} else {
			//更改区间的有边界
			ret[len(ret) - 1][1] = max(ret[len(ret) - 1][1], v[1])
		}
	}
	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}