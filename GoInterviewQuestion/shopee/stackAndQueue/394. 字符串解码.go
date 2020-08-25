package stackAndQueue

import (
	"container/list"
	"strings"
)

func decodeString(s string) string {
	type pair struct {
		multi int
		lastRest string
	}
	//multi用来表示数字
	stack, ret, multi := list.New(), "", 0
	for _, c := range s {
		if c >= '0' && c <= 9 {
			multi = multi*10 + int(c-'0')
		} else if c == '[' {
			stack.PushBack(pair{multi, ret})
			multi, ret = 0, ""
		} else if c == ']' {
			tmp := stack.Back().Value.(pair)
			stack.Remove(stack.Back())
			//字符串拼接
			ret = tmp.lastRest + strings.Repeat(ret, tmp.multi)
		} else {
			ret += string(c)
		}
	}
	return ret
}