package doublePtr

/*
hash表解法
时间复杂度：O(n+m+l)，n，m，l分别为arr1 arr2 arr3的长度
空间复杂度：O(k)。 k为n，m，l中的最大值
*/
func ArraysIntersection1(arr1 []int, arr2 []int, arr3 []int) []int {
	hash := make(map[int]bool)
	for _, v := range arr1 {
		if _, ok := hash[v]; !ok {
			hash[v] = false
		}
	}

	for _, v := range arr2 {
		if _, ok := hash[v]; ok {
			hash[v] = true
		}
	}

	ret := make([]int, 0)
	for _, v := range arr3 {
		if hash[v] == true {
			ret = append(ret, v)
		}
	}
	return ret
}


/*
双指针解法
时间复杂度：O(n)，n为arr1 arr2 arr3总的最小长度
空间复杂度：O(1)。
*/
func ArraysIntersection2(arr1 []int, arr2 []int, arr3 []int) []int {
	ret := make([]int, 0)
	p1, p2, p3 := 0, 0, 0
	// 由于要找的数必须都同时存在于三个数组之中，所以一旦又一个数组遍历完了，就可以确定已经不用找了
	for p1 < len(arr1) && p2 < len(arr2) && p3 < len(arr3) {
		if arr1[p1] == arr2[p2] && arr2[p2] == arr3[p3] {
			ret = append(ret, arr1[p1])
			p1++
			p2++
			p3++
		} else {
			curMax := max(arr1[p1], max(arr2[p2], arr3[p3]))
			if arr1[p1] < curMax {
				p1++
			}
			if arr2[p2] < curMax {
				p2++
			}
			if arr3[p3] < curMax {
				p3++
			}
		}
	}
	return ret
}