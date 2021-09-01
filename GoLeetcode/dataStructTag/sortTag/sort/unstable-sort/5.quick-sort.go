/*
快速排序：不稳定排序，不稳定发生在中枢元素与其他元素交换的时候，5 3 3，中枢元素5与3交换会破坏稳定性
	5 3 3 -> 3 3 3(最右边的一个3赋值给了最左边一个元素) -> 3 3 5
先分大小后递归(分治+递归)，partition时将最左边的元素作为基准，小于等于基准元素的方左边，大于基准元素的放右边

算法时间复杂度
1、最优：O(nlogn); 每一次取到的元素刚好平分整个数组，每一层时间n, 设层为k(第k层有2^k个元素，每一层处理的元素总个数为k)，
 					2^k< n < 2^(k + 1)，k< logn < k + 1, 所以k取logn
2、平均：O(nlogn) ????
3、最差：每一次取到的元素刚好是最小的，O(n^2); T(n) = (n-1) + (n-1) + ... + 1 = n(n-1)/2 = O(n^2)
*/

package unstable_sort

func QuickSort(arr []int, start, end int) {
	if start < end {
		d := partition(arr, start, end)
		QuickSort(arr, start, d-1)
		QuickSort(arr, d+1, end)
	}
}

// 返回基准元素的索引
func partition(arr []int, start, end int) int {
	pivot, i, j := arr[start], start, end
	for i < j {
		// 大于pivot的放右边
		for i < j && arr[j] > pivot  {
			j--
		}
		if i < j {
			// 将小于等于pivot的元素放pivot左边
			arr[i] = arr[j]
		}

		for i < j && arr[i] <= pivot {
			i++
		}
		// 将大于pivot的元素放右边
		if i < j {
			arr[j] = arr[i]
		}
	}
	arr[i] = pivot
	return i
}
