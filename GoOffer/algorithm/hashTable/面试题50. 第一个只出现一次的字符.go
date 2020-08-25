package hashTable


func FirstUniqChar(s string) byte {
	ret := make([]int, 26)
	for _, v := range s {
		ret[v - 'a']++
	}
	for _, v := range s {
		if ret[v - 'a'] == 1 {
			return byte(v)
		}
	}
	return ' '
}