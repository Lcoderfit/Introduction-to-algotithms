package str

import "strings"

func IsFlipedString(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	//如果s1 + s1包含s2, 则s2是s1的字符串轮转，
	//设s1+s1==a1a2a3a4, s2包含在s1+s1中且s2 == a2a3
	//由于len(s1) == len(s2) => len(a1a2) == len(a2a3) => a1 == a3, 同理a2 == a4
	//则a2a3 == a2a1 => 所以s2为s1的字符串轮转
	return strings.Contains(s1 + s1, s2)
}