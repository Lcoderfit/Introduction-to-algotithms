package str


//用一个长度为26的数组累计26个字母出现的次数
func CheckPermutation(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	alphArr := [26]int{0}
	for i := 0; i < len(s1); i++ {
		alphArr[s1[i] - 'a']++
	}
	for i := 0; i < len(s2); i++ {
		if alphArr[s2[i] - 'a'] == 0 {
			return false
		}
		alphArr[s2[i] - 'a']--
	}
	return true
}
