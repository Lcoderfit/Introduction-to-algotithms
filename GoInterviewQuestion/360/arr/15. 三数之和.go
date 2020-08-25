package arr

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	ret := make([][]int, 0)
	for i := 0; i < len(nums) - 2; i++ {
		//nums[i]为三个数中最小的数，如果他>0, 则三个数总和一定大于0
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		p, q := i+1, len(nums) - 1
		for p < q {
			sum := nums[i] + nums[p] + nums[q]
			if sum > 0 {
				q--
			} else if sum < 0 {
				p++
			} else {
				ret = append(ret, []int{nums[i], nums[p], nums[q]})
				//遇到相同元素，右移
				for p < len(nums) - 1 && nums[p] == nums[p+1] {
					p++
				}
				for q > i + 1 && nums[q] == nums[q-1] {
					q--
				}
				p++
				q--
			}
		}
	}
	return ret
}