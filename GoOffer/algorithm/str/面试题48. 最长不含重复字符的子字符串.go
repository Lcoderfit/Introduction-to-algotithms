package str

//方法一：hashTable将重复访问的情况消除
func LengthOfLongestSubstring(s string) int {
	left := 0
	m := make(map[byte]int, 0)
	ret := 0
	for i := 0; i < len(s); i++ {
		if _, ok := m[s[i]]; ok {
			//例如abba这种，第一次重复，left更新为2
			//i为3，s[i] == a,此时为第二次重复,如果是left = m[s[i]] + 1,left的值会倒退为0，
			//所以取left和m[s[i]] + 1的最大值
			left = max(left, m[s[i]] + 1)
		}
		m[s[i]] = i
		ret = max(ret, i-left+1)
	}
	return ret
}

//方法二：256 array
func LengthOfLongestSubstring2(s string) int {
	left, ret := 0, 0
	last := [256]int{}
	//Go语言中初始化数组中所有元素只能遍历赋值
	for i := range last {
		last[i] = -1
	}
	for i := 0; i < len(s); i++ {
		if last[s[i]] > -1 {
			//例如abba这种，第一次重复，left更新为2
			//i为3，s[i] == a,此时为第二次重复,如果是left = m[s[i]] + 1,left的值会倒退为0，
			//所以取left和m[s[i]] + 1的最大值
			left = max(left, last[s[i]]+1)
		}
		last[s[i]] = i
		ret = max(ret, i-left+1)
	}
	return ret
}