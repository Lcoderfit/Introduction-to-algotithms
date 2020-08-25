/*
面试题 10.01. 合并排序的数组.go
时间复杂度：O(m + n)
空间复杂度：O(1)
*/
package sorttag

func Merge(a []int, m int, b []int, n int) {
	pa, pb := m - 1, n - 1
	tail := m + n - 1
	for pa >= 0 || pb >= 0 {
		if pa < 0 {
			a[tail] = b[pb]
			pb--
		} else if pb < 0 {
			a[tail] = a[pa]
			pa--
		} else if a[pa] > b[pb] {
			a[tail] = a[pa]
			pa--
		} else {
			a[tail] = b[pb]
			pb--
		}
		tail--
	}
}