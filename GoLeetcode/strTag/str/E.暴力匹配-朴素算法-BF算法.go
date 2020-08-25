package str

//朴素匹配算法-暴力匹配算法-暴风算法（BF）
//字符串匹配，开始的位置下标都是从0开始的；所以当s[i]和p[j]不匹配时，
//i-j是本次匹配的最开始的下标，i-j+1是下次匹配开始的下标
func BFSearch(haystack string, needle string) int {
	l1 := len(haystack)
	l2 := len(needle)
	i, j := 0, 0
	for i < l1 && j < l2 {
		if haystack[i] == needle[j] {
			i++
			j++
		} else {
			//i往前退j-1格
			i -= j - 1
		}
	}
	if j == l2 {
		return i - j
	}
	return -1
}
