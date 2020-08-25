package dp

import "sort"

//注意，该题一条船最多只能坐两个人
//greedy+双指针
func NumRescueBoats(people []int, limit int) int {
	//先排序，然后双指针从左右两边最轻和最重的人开始
	sort.Ints(people)
	i, j := 0, len(people) - 1
	ret := 0
	//双指针+贪心
	//不用考虑i, j的临界情况，只要想当i==j时，说明还有最后一个人没上船，所以需要再加一条
	for i <= j {
		ret += 1
		if people[i] + people[j] <= limit {
			i++
		}
		j--
	}
	return ret
}