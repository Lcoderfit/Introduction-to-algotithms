/*
选择排序：不稳定排序，例如：5 5 3 当给第一个位置选择元素时，3会与第一个5进行交换，则两个5的相对位置变了(当更小的值在两个相等元素的右边)
给每个位置选择当前未排序元素中的最小值

算法时间复杂度(针对第i个位置，无法避免的需要比较n-i次)
1、最优：O(n^2)
2、平均：O(n^2)
3、最差：O(n^2)
空间复杂度：O(1)
*/
package unstable_sort

// 选择排序
func SelectSort(arr []int) []int {
	if arr == nil || len(arr) == 0 {
		return []int{}
	}
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[i] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}