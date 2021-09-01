package stable_sort

import "math"

/*
	    平均        最好          最坏        空间        稳定性
时间   O(nlogn)   O(nlogn)     O(nlogn)     O(n^2)        稳定
*/
//不需要返回值，直接操作
//先把数组分成两部分，然后对这两部分进行递归分治，最后将排好序的两部分合并起来
//如果len(arr)为奇数，例如为3，mid = 3/2 = 1,
//则前面arr[:mid]=> arr[:1]，前半部分的数组只有一个元素，后半部分有两个；长度为偶数则平分
func MergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	mid := len(arr) / 2

	leftPart := MergeSort(arr[:mid])
	rightPart := MergeSort(arr[mid:])

	merge := make([]int, len(arr))
	i, j, k := 0, 0, 0
	for i < len(leftPart) && j < len(rightPart) {
		if leftPart[i] < rightPart[j] {
			merge[k] = leftPart[i]
			i++
		} else {
			merge[k] = rightPart[j]
			j++
		}
		k++
	}

	for i < len(leftPart) {
		merge[k] = leftPart[i]
		i++
		k++
	}

	for j < len(rightPart) {
		merge[k] = rightPart[j]
		j++
		k++
	}
	return merge
}


func MergeSort1(arr []int) {
	if len(arr) < 2 {
		return
	}
	mid := len(arr) / 2
	// 分治
	MergeSort1(arr[:mid])
	MergeSort1(arr[mid:])

	leftArr := make([]int, mid+1)
	rightArr := make([]int, len(arr)-mid+1)
	//将最后一个元素设置成最大，防止leftArr或者rightArr其中一个比较完之后导致错误
	leftArr[len(leftArr)-1] = math.MaxInt32
	rightArr[len(rightArr)-1] = math.MaxInt32

	//将arr左右两边的数据转移到leftArr和rightArr
	copy(leftArr, arr[:mid])
	copy(rightArr, arr[mid:])

	i, j := 0, 0
	for k := 0; k < len(arr); k++ {
		if leftArr[i] < rightArr[j] {
			arr[k] = leftArr[i]
			i++
		} else {
			arr[k] = rightArr[j]
			j++
		}
	}
}