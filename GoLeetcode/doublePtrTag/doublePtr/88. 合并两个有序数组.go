package doublePtr

/*
时间复杂度：O(m+n)，m是nums1的长度，n是nums2的长度
空间复杂度：O(1)。
*/
//双指针：当nums1长度足够容纳nums1和nums2的长度和的时候；如果是从小到大排序
//则p1指向nums1有数据部分的尾部，p2指向nums2有数据部分的尾部，p指向nums1的(m+n-1)位置
//比较p1与p2指针，将更大的放到p指针位置
func Merge(nums1 []int, m int, nums2 []int, n int) {
	p1, p2, p := m-1, n-1, m+n-1
	for p1 >= 0 && p2 >= 0 {
		if nums1[p1] >= nums2[p2] {
			nums1[p] = nums1[p1]
			p1--
		} else {
			nums1[p] = nums2[p2]
			p2--
		}
		p--
	}
	for p1 >= 0 {
		nums1[p] = nums1[p1]
		p1--
		p--
	}
	//将一个slice的某一部分复制到另一个slice的某一部分
	copy(nums1[:p2+1], nums2[:p2+1])
}
