package randAlgorithm

import "math/rand"

type Solution struct {
	nums []int
}

func Constructor(nums []int) Solution {
	return Solution{nums: nums}
}

func (this *Solution) Pick(target int) int {
	idx := 0
	cnt := 0
	for i, v := range this.nums {
		if v == target {
			cnt += 1

			//this.nums中总共有t个数据与与target相等，需要从这t个数据中随机取出一个，每个取出的概率为1/t
			//蓄水池采样：第一次v == target,以100%的概率取走，后面每次v == target,都以1/cnt的概率取走
			if rand.Intn(cnt) == 0 {
				idx = i
			}
		}
	}
	return idx
}
