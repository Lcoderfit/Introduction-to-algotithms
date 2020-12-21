package binaryFind
/*
1064. 不动点.py
方法1： 二分查找
时间复杂度：O(logn)
空间复杂度：O(1)

case1:
0 1 2 3 4

case2:
-10 -5 0 3 7
 */

func FixedPoint(A []int) int {
	if len(A) == 0 {
		return -1
	}
	i, j := 0, len(A) - 1
	for i < j {
		mid := (i + j) / 2
		if A[mid] >= mid {
			j = mid
		} else {
			i = mid + 1
		}
	}
	if A[i] == i {
		return i
	}
	return -1
}