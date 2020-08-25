package arr


func MaxArea(height []int) int {
	i, j := 0, len(height) - 1
	ret := 0
	//双指针
	for i < j {
		//木桶原理，水的高度由短的木板决定
		//比较左右两边的木板，移动短板指针
		if height[i] < height[j] {
			ret = max(ret, (j-i)*height[i])
			i++
		} else {
			ret = max(ret, (j-i)*height[j])
			j--
		}
	}
	return ret
}
