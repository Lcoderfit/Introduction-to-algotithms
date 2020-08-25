package doublePtr


/*
双指针解法：该算法的一个巧妙之处在于，当判断i, j所指向的字符是否为元音时，通过三个if即可
一般是通过四个if和else if等来判断的：
if isVowei(sByte[i]) && isVowei(sByte[j])
else if !isVowei(sByte[i])
else if !isVowei(sByte[j])
else

本例通过两个！和一个正判断的&&简化了逻辑

时间复杂度：O(n)，n是string的长度
空间复杂度：O(1)。
*/
func ReverseVowels1(s string) string {
	isVowei := func(b byte) bool {
		return b == 'A' || b == 'a' || b == 'E' || b == 'e' ||b == 'I' ||
			b == 'i' || b == 'O' || b == 'o' || b == 'U' || b == 'u'
	}
	i, j := 0, len(s) - 1
	sByte := []byte(s)
	for i < j {
		if !isVowei(sByte[i]) {
			i++
		}
		if !isVowei(sByte[j]) {
			j--
		}
		if isVowei(sByte[i]) && isVowei(sByte[j]) {
			sByte[i], sByte[j] = sByte[j], sByte[i]
			i++
			j--
		}
	}
	return string(sByte)
}