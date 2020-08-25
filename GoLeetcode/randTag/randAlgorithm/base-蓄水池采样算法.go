package randAlgorithm

import (
	"math/rand"
)

//从n个元素中随机选取k个
func RandKFromN(nums []int, k int) []int {
	ret := make([]int, k)
	copy(ret, nums[:k])
	for i := k; i < len(nums); i++ {
		//以k/i的概率替换
		if rand.Intn(i) < k {
			ret[rand.Intn(k)] = nums[i]
		}
	}
	return ret
}
