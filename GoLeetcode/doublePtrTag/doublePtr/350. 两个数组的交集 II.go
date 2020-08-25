package doublePtr

import "sort"

/*
hash解法
时间复杂度：O(m+n)，m是nums1的长度，n是nums2的长度
空间复杂度：O(min(m, n))。
*/
func Intersect1(nums1 []int, nums2 []int) []int {
    hash := make(map[int]int)
    for _, v := range nums1 {
        hash[v]++
    }
    ret := make([]int, 0)
    for _, v := range nums2 {
        if hash[v] > 0 {
            hash[v]--
            ret = append(ret, v)
        }
    }
    return ret
}


/*
双指针解法
时间复杂度：O(nlogn + mlogm)，m是nums1的长度，n是nums2的长度
空间复杂度：O(1)。
*/
func Intersect2(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	i, j, k := 0, 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			i++
		} else if nums1[i] > nums2[j] {
			j++
		} else {
			nums1[k] = nums2[j]
			k++
			i++
			j++
		}
	}
	return nums1[:k]
}