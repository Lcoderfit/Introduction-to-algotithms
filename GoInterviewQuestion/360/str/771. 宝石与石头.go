package str


func numJewelsInStones(J string, S string) int {
	var sArr [128]int
	//先统计S中各种石头的数量
	for _, v := range S {
		sArr[v]++
	}
	//然后判断有多少是宝石
	ret := 0
	for _, v := range J {
		ret += sArr[v]
	}
	return ret
}