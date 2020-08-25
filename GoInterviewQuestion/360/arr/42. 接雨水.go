package arr


//最优解
//找中间最大的一根柱子可以简化成双指针
/*
//i, j一定同时指向最大的那根柱子
func test(a []int) {
    i, j := 0, len(a) - 1
    for i < j {
        if a[i] < a[j] {
            i++
        } else {
            j--
        }
    }
}
*/

func trap(height []int) int {
	if len(height) <= 2 {
		return 0
	}
	ret, leftMax, rightMax := 0, 0, 0
	for left, right := 0, len(height) - 1; left < right; {
		//此时left指针在最大柱子的左边
		if height[left] < height[right] {
			//以最大值为分界线，计算左右两边的积水量
			//以左边为例，设当前柱子高度为leftMax，如果他的下一个柱子高度小于leftMax，则一定有积水
			//设leftMax的位置为i，则一直遍历i+1....j, 直到找到height[j] >= leftMax, 此时不会产生积水
			//更新leftMax的值为height[j]
			if leftMax <= height[left] {
				leftMax = height[left]
			} else {
				ret += leftMax - height[left]
			}
			left++
			//此时right指针在最大值的右边
		} else {
			if rightMax <= height[right] {
				rightMax = height[right]
			} else {
				ret += rightMax - height[right]
			}
			right--
		}
	}
	return ret
}

func trap1(height []int) int {
	if len(height) <= 2 {
		return 0
	}
	m := 0
	//先找出数组中的最大值的下标
	for i, v := range height {
		if height[m] < v {
			m = i
		}
	}
	//以最大值为分界线，计算左右两边的积水量
	//以左边为例，设当前柱子高度为cur，如果他的下一个柱子高度小于cur，则一定有积水
	//设cur的位置为i，则一直遍历i+1....j, 知道找到height[j] > cur, 此时不会产生积水
	//更新cur的值为height[j]
	ret := 0
	cur := height[0]
	for i := 1; i <= m; i++ {
		if cur > height[i] {
			ret += cur - height[i]
		} else {
			cur = height[i]
		}
	}
	cur = height[len(height) - 1]
	for j := len(height) - 1; j >= m; j-- {
		if cur > height[j] {
			ret += cur - height[j]
		} else {
			cur = height[j]
		}
	}
	return ret
}