package str


//KMP字符串匹配算法
/*
主串长度为M，子串长度为N
时：O(m+n)
空：O(n)
 */
func KmpSearch(haystack string, needle string, next []int) int {
	l1 := len(haystack)
	l2 := len(needle)
	i, j := 0, 0
	for i < l1 && j < l2 {
		//j == -1是用来防止陷入无限循环的
		//若next[0] == 0; 则当第一个字符不匹配时(j==0)，也就是haystack[i] != needle[0], 则j = next[0] = 0
		//那么下次匹配仍然是haystack[i]与needle[0]进行匹配，就会陷入无限循环

		//若next[0] == -1; 则当第一个字符不匹配时(j==0)，也就是haystack[i] != needle[0], 则j = next[0] = -1
		//那么下次匹配时j == -1为True，i++, j++(j变为0)，下次匹配就是haystack[i](此时的i自增了1)与needle[0]比较了
		//不会陷入无限循环
		if j == -1 || haystack[i] == needle[j] {
			i++
			j++
		} else {
			//next[j]表示j位置不匹配时下一次开始匹配的位置
			j = next[j]
		}
	}
	if j == l2 {
		return i - j
	}
	return -1
}