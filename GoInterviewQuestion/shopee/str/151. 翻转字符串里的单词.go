package str

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	s1 := strings.Split(strings.TrimSpace(s), " ")
	for i := 0; i < len(s1); i++ {
		if s1[i] == "" {
			s1 = append(s1[:i], s1[i+1:]...)
			i--
		}
	}
	fmt.Println(s1, len(s1))
	ret := ""
	for i := len(s1) - 1; i >= 0; i-- {
		if i == len(s1) - 1 {
			ret += s1[i]
		} else {
			ret += " " + s1[i]
		}
	}
	return ret
}