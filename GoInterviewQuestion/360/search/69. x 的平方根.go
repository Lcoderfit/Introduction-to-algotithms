package search


func mySqrt(x int) int {
	if x == 1 {
		return 1
	}
	//二分查找的思想
	i, j := 1, x/2
	for i <= j {
		mid := (i+j)/2
		if mid*mid > x {
			j = mid - 1
		} else if mid * mid < x {
			i = mid + 1
		} else {
			return mid
		}
	}
	return i - 1
}

//i < j的解法
func mySqrt1(x int) int {
	if x <= 1 {
		return x
	}
	//二分查找的思想
	i, j := 1, x/2
	for i < j {
		//i = mid, 就需要中间值偏向j
		mid := (i+j+1)/2
		if mid*mid > x {
			j = mid - 1
		} else if mid * mid < x {
			i = mid
		} else {
			return mid
		}
	}
	return i
}