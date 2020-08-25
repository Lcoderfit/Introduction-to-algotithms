package doublePtr


/*
双指针+排序解法
时间复杂度：O(nlogn)，n为A的长度
空间复杂度：O(n)。
*/
func SortedSquares1(A []int) []int {
	ret := make([]int, len(A))
	s, e := 0, len(A)-1
	idx := len(A) - 1
	for s <= e {
		if A[s]*A[s] > A[e]*A[e] {
			ret[idx] = A[s] * A[s]
			s++
		} else {
			ret[idx] = A[e] * A[e]
			e--
		}
		idx--
	}

	return ret
}


/*
双指针解法
时间复杂度：O(n)，n为A的长度
空间复杂度：O(n)。
*/
func SortedSquares2(A []int) []int {
	i, j := -1, -1
	for k, v := range A {
		if A[k] < 0 {
			A[k] = v*v
			i = k
			continue
		}
		j = k
		break
	}
	ret := make([]int, len(A))
	k := 0
	for i >= 0 && j >= 0 && j < len(A) {
		v := A[j]*A[j]
		if A[i] < v {
			ret[k] = A[i]
			k++
			i--
		} else {
			ret[k] = v
			j++
			k++
		}
	}

	for i >= 0 {
		ret[k] = A[i]
		k++
		i--
	}
	for j >= 0 && j < len(A) {
		ret[k] = A[j]*A[j]
		k++
		j++
	}
	return ret
}