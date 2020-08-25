package doublePtr

import "sort"

/*
双指针解法
时间复杂度：O(nlogn + mlogm)，m是nums1的长度，n是nums2的长度
空间复杂度：O(1)。
*/
func Intersection(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	i, j := 0, 0
	ret := make([]int, 0)
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			i++
		} else if nums1[i] > nums2[j] {
			j++
		} else {
			for i < len(nums1) - 1 && nums1[i] == nums1[i+1] {
				i++
			}
			for j < len(nums2) - 1 && nums2[j] == nums2[j+1] {
				j++
			}
			ret = append(ret, nums1[i])
			i++
			j++
		}
	}
	return ret
}


/*
hash表解法: 先遍历一个表，如果存在某元素则将hash表映射为true
然后遍历另外一个表，如果存在，则将元素放入ret中，然后将hash映射为false（避免重复）

时间复杂度：O(m+n)，m是nums1的长度，n是nums2的长度
空间复杂度：O(m)。
*/
func Intersection1(nums1 []int, nums2 []int) []int {
	ret := make([]int, 0)
	hash := make(map[int]bool)

	for _, v := range nums1 {
		hash[v] = true
	}
	for _, v := range nums2 {
		if hash[v] == true {
			ret = append(ret, v)
			hash[v] = false
		}
	}
	return ret
}