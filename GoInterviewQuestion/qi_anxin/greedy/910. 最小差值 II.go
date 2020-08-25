package greedy

import "sort"

func smallestRangeII(A []int, K int) int {
	if len(A) == 1 {
		return 0
	}
	sort.Ints(A)
	//不对slice进行二分（前一半上移，后一半下移，这里称为二分）
	ret := A[len(A) - 1] - A[0]
	for i := 0; i < len(A) - 1; i++ {
		//二分移动，前一半全+K, 后一半全减K
		//二分移动后，最大值为max(前一半的最大值， 后一半的最大值)
		high := max(A[i]+K, A[len(A) - 1]-K)
		low := min(A[0]+K, A[i+1]-K)
		ret = min(ret, high-low)
	}
	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//匿名函数
func SmallestRangeII(A []int, K int) int {
	sort.Ints(A)
	l := len(A)
	if l <= 1 {
		return 0
	}

	r := A[l-1] - A[0]

	for i := 0; i < l-1; i++ {
		min := func(m, n int) int {
			if m > n {
				return n
			}
			return m
		}(A[0]+K, A[i+1]-K)

		max := func(m, n int) int {
			if m > n {
				return m
			}
			return n
		}(A[i]+K, A[l-1]-K)

		r = func(m, n, rcur int) int {
			if rcur > m-n {
				return m - n
			}
			return rcur
		}(max, min, r)
	}

	return r
}