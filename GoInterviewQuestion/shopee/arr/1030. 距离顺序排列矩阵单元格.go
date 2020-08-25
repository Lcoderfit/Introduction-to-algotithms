package arr

import "sort"

//用hashMap存储dist到作标的映射
func AllCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {
	ret := [][]int{}
	//用于存储相同距离的位置
	distMap := make(map[int][][]int)
	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			//先求出距离dist, 再存储dist对应的位置
			dist := abs(i - r0) + abs(j - c0)
			distMap[dist] = append(distMap[dist], []int{i, j})
		}
	}

	keySort := make([]int, 0)
	for k, _ := range distMap {
		keySort = append(keySort, k)
	}
	//将所有的key进行排序
	sort.Ints(keySort)

	//根据dist从小到大放入位置
	for _, dist := range keySort {
		for _, v := range distMap[dist] {
			ret = append(ret, v)
		}
	}

	return ret
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}