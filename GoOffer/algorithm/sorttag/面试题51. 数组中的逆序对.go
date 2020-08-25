package sorttag

import (
	"math"
)

//归并排序解法：O(NlogN)
func ReversePairs(nums []int) int {
	ret := 0
	if len(nums) < 2 {
		return ret
	}
	mid := len(nums)/2
	ret += ReversePairs(nums[:mid]) + ReversePairs(nums[mid:])
	leftArr, rightArr := make([]int, mid), make([]int, len(nums)-mid)

	copy(leftArr, nums[:mid])
	copy(rightArr, nums[mid:])
	i, j, k := 0, 0, 0
	for i < mid && j < len(nums) - mid {
		if leftArr[i] <= rightArr[j] {
			nums[k] = leftArr[i]
			i++
			//此时leftArr[i] > rightArr[0...j-1]
			ret += j
		} else {
			nums[k] = rightArr[j]
			j++
		}
		k++
	}
	for i < mid {
		nums[k] = leftArr[i]
		i++
		k++
		//此时leftArr[i] > rightArr中所有元素
		ret += j
	}
	for j < mid {
		nums[k] = rightArr[j]
		k++
		j++
	}
	return ret
}

//二：归并优化
func ReversePairs2(nums []int) int {
	ret := 0
	if len(nums) < 2 {
		return ret
	}
	mid := len(nums)/2
	ret += ReversePairs2(nums[:mid]) + ReversePairs2(nums[mid:])
	leftArr, rightArr := make([]int, mid+1), make([]int, len(nums)-mid+1)
	//初始化最后一个元素为最大
	leftArr[len(leftArr) - 1], rightArr[len(rightArr) - 1] = math.MaxInt32, math.MaxInt32

	copy(leftArr, nums[:mid])
	copy(rightArr, nums[mid:])

	i, j := 0, 0
	for k := 0; k < len(nums); k++ {
		if leftArr[i] <= rightArr[j] {
			nums[k] = leftArr[i]
			i++
			//此时leftArr[i] > rightArr[0...j-1]
			ret += j
		} else {
			nums[k] = rightArr[j]
			j++
		}
	}
	return ret
}