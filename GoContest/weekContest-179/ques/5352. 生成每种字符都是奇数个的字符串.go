package ques


func GenerateTheString(n int) string {
	a, b := "a", "b"
	ret := ""
	if n % 2 == 1 {
		ret = addStr(ret, a, n)
	} else {
		ret = addStr(ret, a, n - 1)
		ret +=  b
	}
	return ret
}

func addStr(ret, s string, k int) string {
	for i := 0; i < k; i++ {
		ret += s
	}
	return ret
}
