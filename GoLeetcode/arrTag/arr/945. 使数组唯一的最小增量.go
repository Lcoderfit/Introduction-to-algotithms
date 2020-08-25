package arr

import "sort"

//一：计数法
func MinIncrementForUnique(A []int) int {
	if len(A) == 0 {
		return 0
	}
	//该题不变性在于数据范围可以确定，0 <= A.length <= 40000
	//如果所有40000个数据的值都为40000（0 <= A[i] < 40000）
	//则需要[40001, 40002...., 79999]的数据来进行替换,所以经过操作之后的最大值<80000
	count := make([]int, 80000)
	//利用hashtable统计重复的数字及重复次数
	for _, v := range A {
		count[v]++
	}

	//moveNum为需要进行move操作的数的总个数
	moveNum, ret := 0, 0
	for i := 0; i < 80000; i++ {
		if count[i] > 1 {
			//count[i] - 1为重复的数组的个数
			moveNum += count[i] - 1
			//减去所有重复数字的总和
			ret -= i*(count[i] - 1)
		} else if moveNum > 0 && count[i] == 0 {
			moveNum--
			//加上经过move操作后的数据
			ret += i
		}
	}
	return ret
}

//二：排序
func MinIncrementForUnique2(A []int) int {
	ret := 0
	sort.Ints(A)
	// 遍历数组，若当前元素小于等于它的前一个元素，则将其变为前一个数+1
	for i := 1; i < len(A); i++ {
		if A[i] <= A[i-1] {
			tmp := A[i]
			A[i] = A[i-1]+1
			ret += A[i] - tmp
		}
	}
	return ret
}

//三：计数排序
func MinIncrementForUnique3(A []int) int {
	// counter数组统计每个数字的个数。
	//（这里为了防止下面遍历counter的时候每次都走到40000，所以设置了一个max，这个数据量不设也行，再额外设置min也行）
	count := make([]int, 40001)
	minVal, maxVal := 40001, 0
	for _, v := range A {
		count[v]++
		maxVal = max(maxVal, v)
		minVal = min(minVal, v)
	}
	// 遍历counter数组，若当前数字的个数cnt大于1个，则只留下1个，其他的cnt-1个后移
	ret := 0
	for i := minVal; i < maxVal; i++ {
		if count[i] > 1 {
			moveTimes := count[i] - 1
			ret += moveTimes
			count[i] = 1
			count[i+1] += moveTimes
		}
	}
	// 最后, counter[max+1]里可能会有从counter[max]后移过来的，counter[max+1]里只留下1个，其它的d个后移。
	// 设 max+1 = x，那么后面的d个数就是[x+1,x+2,x+3,...,x+d],
	// 因此操作次数是[1,2,3,...,d],用求和公式求和。
	moveTimes := count[maxVal] - 1
	ret += moveTimes*(moveTimes+1)/2
	return ret
}