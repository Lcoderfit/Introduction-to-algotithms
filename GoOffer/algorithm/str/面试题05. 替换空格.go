package str


func ReplaceSpace(s string) string {
	ret := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			ret = append(ret, s[i])
		} else {
			ret = append(ret, []byte("%20")...)
		}
	}
	return string(ret)
}