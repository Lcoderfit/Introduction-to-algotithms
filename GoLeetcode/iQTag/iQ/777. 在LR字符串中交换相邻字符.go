package iQ

import (
	"strings"
)

/*
a := "LXXRRXX"
s := strings.Replace(a, "X", "", -1)
fmt.Println(s)
 */
func CanTransform(start string, end string) bool {
	if strings.Replace(start, "X", "", -1) != strings.Replace(end, "X", "", -1) {
		return false
	}

	j := 0
	for i := 0; i < len(start); i++ {
		if start[i] == 'L' {
			for end[j] != 'L' {
				j++
			}
			if i < j {
				return false
			}
			j++
		}
	}

	j = 0
	for i := 0; i < len(end); i++ {
		if start[i] == 'R' {
			for end[j] != 'R' {
				j++
			}
			if i > j {
				return false
			}
			j++
		}
	}
	return true
}
