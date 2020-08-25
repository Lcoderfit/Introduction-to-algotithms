package main

import (
	"GoLeetcode/strTag/str"
	"fmt"
)

func main() {

	//E.暴力匹配-朴素算法-BF算法.go
	bfSearch()

	//409. 最长回文串.go
	//longestPalindrome()

	//438. 找到字符串中所有字母异位词.go
	//findAnagrams()

	//1071. 字符串的最大公因子.go
	//gcdOfStrings1()
}

//E.暴力匹配-朴素算法-BF算法.go
func bfSearch() {
	var haystack, needle string
	for {
		fmt.Scanln(&haystack, &needle)
		ret := str.BFSearch(haystack, needle)
		fmt.Println("E.暴力匹配-朴素算法-BF算法.go: ", ret)
	}
}

//409. 最长回文串.go
func longestPalindrome() {
	var s string
	for {
		//需要传入指针作为参数
		fmt.Scanln(&s)
		ret := str.LongestPalindrome(s)
		fmt.Println("longestPalindrome: ", ret)
	}
}

//438. 找到字符串中所有字母异位词.go
func findAnagrams() {
	var s, p string
	for {
		fmt.Scanln(&s)
		fmt.Scanln(&p)
		//cbaebabacd
		//abc
		ret := str.FindAnagrams(s, p)
		fmt.Println("findAnagrams: ", ret)
	}
}

//1071. 字符串的最大公因子.go
func gcdOfStrings1() {
	var str1, str2 string
	for {
		fmt.Scanf("%s%s", &str1, &str2)
		ret := str.GcdOfStrings1(str1, str2)
		fmt.Println("gcdOfString1: ", ret)
	}
}