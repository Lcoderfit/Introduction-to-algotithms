package str

var m = make(map[string]bool, 0)

func Permutation(s string) []string {
	if len(s) == 0 {
		return []string{}
	}
	if len(s) <= 1 {
		return []string{s}
	}
	m := make(map[string]bool)
	ret := make([]string, 0)
	for i := 0; i < len(s); i++ {
		t := s[i]
		tailArr := Permutation(s[:i] + s[i+1:])
		for _, str := range tailArr {
			k := string(t) + str
			if m[k] == false {
				m[k] = true
				ret = append(ret, k)
			}
		}
	}
	return ret
}