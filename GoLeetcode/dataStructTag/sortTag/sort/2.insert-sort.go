/*

插入排序与冒泡排序的区别，冒泡排序是先确定最右边的元素；而插入排序是先确定最左边的元素

插入排序：稳定排序，当a[index - 1] > a[index]时才交换，相等时不交互
算法时间复杂度
1、最优：O(n), 数组恰好有序，插入每个元素时都只比一次
2、平均：O(n^2)
3、最差：O(n^2), 第i个元素需要比较i - 1次，1 + 2 + ... + n - 1
*/
package sort

func InsertSort(arr []int) []int {
	if arr == nil || len(arr) == 0 {
		return []int{}
	}
	for i := 1; i < len(arr); i++ {
		// 记录当前需要排序元素的值
		k := arr[i]
		j := i - 1
		for ; j >= 0 && arr[j] > k; j-- {
			arr[j+1] = arr[j]
		}
		arr[j+1] = k
	}
	return arr
}
