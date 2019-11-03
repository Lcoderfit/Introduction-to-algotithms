package main

import (
	"bytes"
	"fmt"
	"strings"
	"unicode"
)

func main() {
	a := "Lcoder"
	b := "lcodeR"

	fmt.Println(unicode.IsDigit('k'), strings.ToUpper(a), strings.ToLower(b))
	for _, i := range a {
		fmt.Println(unicode.IsLetter(i), unicode.IsUpper(rune(a[0])))
	}
	fmt.Printf("%q", a[0])

	p := "file/fuck/example.go"
	fmt.Println(basename(p), basename1(p))
	fmt.Println(comma("123456789"))

	k := "abc"
	g := []byte(k)
	fmt.Println(g)
	for i := 0; i < len(g); i++ {
		fmt.Println(g[i])
	}
	fmt.Println('k')

	fmt.Println(intsToString([]int{1, 2, 3}))

	y := "lcoder"
	fmt.Fprintf(&y, "%s", "fit")
	fmt.Println(y)
}

func basename(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

func basename1(s string) string {
	slash := strings.LastIndex(s, "/") //如果没有则返回-1
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}
