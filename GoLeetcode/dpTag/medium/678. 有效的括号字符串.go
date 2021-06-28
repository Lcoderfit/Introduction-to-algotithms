package medium

func CheckValidString(s string) bool {
	l, h := 0, 0
	for _, c := range s {
		if c == '(' {
			l++
			h++
		} else if c == '*' {
			if l > 0 {
				l--
			}
			h++
		} else {
			if l > 0 {
				l--
			}
			h--
		}
		if h < 0 {
			return false
		}
	}
	return l == 0
}
