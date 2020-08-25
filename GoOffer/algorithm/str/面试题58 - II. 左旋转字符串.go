package str


func ReverseLeftWords(s string, n int) string {
	if n >= len(s) {
		return s
	}
	return string(s[n:]+s[:n])
}