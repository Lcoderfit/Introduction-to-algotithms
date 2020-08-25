package math


func reachNumber(target int) int {
	//逐步计算所能到达的最远距离L
	//当target < L时，说明，步数不够，需要继续加步数
	//统一当正向的值来处理，因为从0到target和从0到-target的步数是相同的
	if target < 0 {
		target = -target
	}
	maxDist := 0
	stepCount := 0
	//该题规律，假设k步所能到达的最大距离是L
	//那么k步所能到的点是L, L-2, L-4 ....，如果L为奇数，其它所能达到的点也同为奇数
	//target的奇偶性必须与L相同，否则K步所能到达的点不可能包含target
	for maxDist < target || maxDist%2 != target%2 {
		stepCount++
		maxDist += stepCount
	}
	return stepCount
}