package arr

import "sort"

//方法一：摩尔投票法
func MajorityElement(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	votes, cur := 0, 0
	for _, v := range nums {
		if votes == 0 {
			cur = v
		}
		if cur == v {
			votes++
		} else {
			votes--
		}
	}
	return cur
}

//方法二：排序
func majorityElement(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}