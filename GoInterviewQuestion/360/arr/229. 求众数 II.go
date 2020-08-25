package arr


func majorityElement(nums []int) []int {
	//出现次数超过1/3，说明最多有2个候选人
	//定义2个候选人变量，2个投票变量分别表示候选人的票数
	p1, p2 := 0, 0
	votes1, votes2 := 0, 0
	//投票阶段
	for _, v := range nums {
		//判断是否相等的必须放前面，这样避免了p1和p2相等的情况
		if p1 == v {
			votes1++
			continue
		}
		if p2 == v {
			votes2++
			continue
		}
		if votes1 == 0 {
			p1 = v
			votes1++
			continue
		}
		if votes2 == 0 {
			p2 = v
			votes2++
			continue
		}
		//v与候选人1和2都不相同，票数减少
		votes1--
		votes2--
	}
	//计数阶段
	votes1, votes2 = 0, 0
	for _, v := range nums {
		if p1 == v {
			votes1++
		} else if p2 == v {
			votes2++
		}
	}
	ret := make([]int, 0)
	if votes1 > len(nums)/3 {
		ret = append(ret, p1)
	}
	if votes2 > len(nums)/3 {
		ret = append(ret, p2)
	}
	return ret
}