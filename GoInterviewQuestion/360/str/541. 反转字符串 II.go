package str


//最优版
func ReverseStr(s string, k int) string {
	sByte := []byte(s)
	for i := 0; i*k < len(s); i += 2 {
		start, end := i*k, i*k+k-1
		if end >= len(s) {
			end = len(s) - 1
		}
		for ; start < end; start, end = start+1, end-1 {
			sByte[start], sByte[end] = sByte[end], sByte[start]
		}
	}
	return string(sByte)
}


func reverseStr(s string, k int) string {
	sByte := []byte(s)
	var i int
	for i = 0; i+k-1 < len(s); i += 2*k {
		reverse(sByte, i, i+k-1)
	}
	if i < len(sByte) {
		reverse(sByte, i, len(sByte) - 1)
	}
	return string(sByte)
}

func reverse(sByte []byte, i, j int) {
	for ; i < j; i, j = i+1, j-1 {
		sByte[i], sByte[j] = sByte[j], sByte[i]
	}
}