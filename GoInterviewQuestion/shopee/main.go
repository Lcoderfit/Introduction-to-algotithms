package main

import (
	"GoInterviewQuestion/shopee/dp"
	"fmt"
	"strconv"
	"strings"
)

//
//import (
//	"fmt"
//)
//
//func main() {
//	a := "abcdedfs"
//	fmt.Printf("%T\n", a[0])
//	for _, v := range a {
//		fmt.Printf("%T\n", v)
//		break
//	}
//	for i := 0; i < len(a); i++ {
//		fmt.Printf("%T\n", a[i])
//		fmt.Println()
//		break
//	}
//
//	if a[0] == 'a' {
//		fmt.Println(true)
//	}
//	fmt.Printf("%T\n", 'a')
//	b := make([]byte, 0)
//	b = append(b, a[0])
//	fmt.Println(b)
//	b = append(b, 'a')
//	fmt.Println(b)
//
//	//切片左边界最大允许等于切片的长度
//	e := []int{1, 2}
//	fmt.Println(e[2:])
//	//切片右边界最大不超过切片的长度
//	fmt.Println(e[:2])
//	//fmt.Println(e[:3]) 这句是错误的
//
//	f, g := "20", "2"
//	fmt.Println("f < g: ", f < g)
//
//	var m map[int]int
//	m = make(map[int]int)
//	m[0] = 1
//	fmt.Println(m)
//
//	n := make(map[int]int)
//	n[0] = 0
//
//	//k, _ := json.Marshal(n)
//	//fmt.Println(k)
//	//p := make(map[int]int)
//	//_ = json.Unmarshal(k, &p)
//	//fmt.Println(p)
//}

func main() {
	//三数之和为0
	//threeSum()

	//不同走法
	diffWay()
}

//不同走法
func diffWay() {
	//不同走法
	var m, n int
	fmt.Scanln(&m)
	fmt.Scanln(&n)
	ret := dp.CountPathNum(m, n)
	fmt.Println(ret)
}

//三数之和
func threeSum() {
	//找组合
	var s string
	fmt.Scanf("%s", &s)
	array := make([]int, 0)
	for _, v := range strings.Split(s, ",") {
		t, _ := strconv.Atoi(v)
		array = append(array, t)
	}
	ret := arr.FindCombine(array)
	fmt.Println(ret)
	for i := 0; i < len(ret); i++ {
		if i == 0 {
			//fmt.Printf("%v", ret[i])
			fmt.Printf("[")
			for j := 0; j < len(ret[i]); j++ {
				if j == 0 {
					fmt.Printf("%d", ret[i][j])
				} else {
					fmt.Printf(",%d", ret[i][j])
				}
			}
			fmt.Printf("]")
		} else {
			//fmt.Printf(",%v", ret[i])
			fmt.Printf(",[")
			for j := 0; j < len(ret[i]); j++ {
				if j == 0 {
					fmt.Printf("%d", ret[i][j])
				} else {
					fmt.Printf(",%d", ret[i][j])
				}
			}
			fmt.Printf("]")
		}
	}
}