/*
冒泡排序：稳定排序，当a[j] > a[j + 1]时才交换，相等时不交互
外层循环为对比的轮次数，一共n-1次（因为排好了n-1个之后剩下一个自然也就排好了）
内层循环每次从索引0开始，对比arr[j]与arr[j+1]的大小，从小到大排序就将大的往后移，从大到小排序就将小的往后移
下一轮对比的次数比上一轮要少一次

优化：添加一个swapped变量来表示一轮的交换中，是否存在一个元素都没交换的情况，
	如果一个元素都没交换，则表示所有元素已经有序了，直接退出外外层循环

算法时间复杂度
1、最优：冒泡排序优化一的最优情况是O(n)
2、平均：有n-1种情况（只比较一轮结束，比较两轮结束。。。比较n-1轮结束），
	每种情况需要比较的次数为i*(n - 1)， (1+2+...+n-1)*(n-1)/(n-1) = O(n^2)
3、最差：O(n^2) 比较n-1轮，每一轮n-i-1次


周末复习的时候写成 (s *Sort) xxxxSort的形式
*/
package sort

func BubbleSort(arr []int) []int {
	if arr == nil || len(arr) == 0 {
		return []int{}
	}
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

// 冒泡排序优化，面试直接写这个
func BubbleSort1(arr []int) []int {
	if arr == nil || len(arr) == 0 {
		return []int{}
	}
	for i := 0; i < len(arr)-1; i++ {
		swapped := false
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				swapped = true
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
		if !swapped {
			break
		}
	}
	return arr
}
