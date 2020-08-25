package arr

import "sort"

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	var p, q int
	ret := nums[0] + nums[1] + nums[2]
	for i := 0; i < len(nums)-2; i++ {
		p, q = i+1, len(nums) - 1
		for p < q {
			sum := nums[i] + nums[p] + nums[q]
			if sum == target {
				return sum
			} else if sum > target {
				q--
			} else {
				p++
			}
			if abs(sum-target) < abs(ret-target) {
				ret = sum
			}
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