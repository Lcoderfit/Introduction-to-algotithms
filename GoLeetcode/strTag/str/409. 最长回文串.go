package str

//hashtable
func LongestPalindrome(s string) int {
	ret := 0
	//hashMap
	// m := make(map[byte]int, 0)

	//ascii array, 字符串只包含大小写字母，都可以考虑用ascii array进行处理
	//'z' => 97+26-1==122
	m := [128]int{}
	for i := range s {
		m[s[i]]++
	}
	flag := false
	for _, v := range m {
		if !flag && v%2 == 1 {
			flag = true
		}
		ret += v/2*2
	}
	if flag {
		ret++
	}
	return ret
}