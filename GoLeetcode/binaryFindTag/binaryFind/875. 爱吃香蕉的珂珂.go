package binaryFind

func MinEatingSpeed(piles []int, H int) int {
	//二分查找算法时间复杂度为logN
	maxSpeed := 0
	for _, v := range piles {
		if maxSpeed < v {
			maxSpeed = v
		}
	}

	i, j := 1, maxSpeed
	//循环条件为i<j, 则当i == j时，即找到了满足条件的最小速度k
	for i < j {
		mid := (i+j)/2
		count := 0
		for _, v := range piles {
			//向上取整
			count += (v+mid-1)/mid
			if count > H {
				break
			}
		}
		//二分缩小范围
		//需要花更多时间，说明速度小了，必须增大i
		if count > H {
			i = mid + 1
		} else {
			//如果count == H, 则此时的mid有可能就是答案，所以j = mid而不是j = mid -1
			j = mid
		}
	}
	return i
}