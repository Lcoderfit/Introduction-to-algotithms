package str

import "strconv"

//方法一：用[]byte
func CompressString(S string) string {
	if len(S) == 0 {
		return S
	}
	count := 1
	ret := make([]byte, 0)
	for i := 1; i < len(S); i++ {
		if S[i] != S[i-1] {
			ret = append(ret, S[i-1])
			ret = append(ret,  []byte(strconv.Itoa(count))...)
			count = 1
		} else {
			count++
		}
	}
	ret = append(ret, S[len(S) - 1])
	ret = append(ret,  []byte(strconv.Itoa(count))...)
	if len(ret) >= len(S) {
		return S
	}
	return string(ret)
}

//方法二：string
func CompressString1(S string) string {
	if len(S) == 0 {
		return S
	}
	count := 1
	ret := ""
	for i := 1; i < len(S); i++ {
		if S[i] != S[i-1] {
			ret += string(S[i-1]) + strconv.Itoa(count)
			count = 1
		} else {
			count++
		}
	}
	ret += string(S[len(S) - 1]) + strconv.Itoa(count)
	if len(ret) >= len(S) {
		return S
	}
	return string(ret)
}