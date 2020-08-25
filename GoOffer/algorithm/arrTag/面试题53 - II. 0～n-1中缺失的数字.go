package arrTag



//方法一
/*
平均算法时间复杂度下降到：O(logn)
最坏情况仍然是O(n)
 */
func MissingNumber(nums []int) int {
	i, j := 0, len(nums)
	var mid int
	for i < j {
		mid = (i+j)/2
		if nums[mid] == mid {
			i = mid + 1
		} else {
			j = mid
		}
	}
	return i
}

//方法二
func MissingNumber1(nums []int) int {
	for i, v := range nums {
		if i != v {
			return i
		}
	}
	return len(nums)
}

//方法三
func MissingNumber2(nums []int) int {
	vSum, iSum := 0, 0
	for i, v := range nums {
		vSum += v
		iSum += i+1
	}
	return iSum - vSum
}