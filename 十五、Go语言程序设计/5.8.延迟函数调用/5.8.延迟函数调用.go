package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path"

	// "strings"
	"net/http"
	"sync"
	"time"
	// "golang.org/x/net/html"
)

func main() {
	fmt.Println("f1 result:", f1())
	fmt.Println("f2 result:", f2())

	bigSlowOperation()

	p := a(4)
	fmt.Println("a(4)", p)

	fmt.Println("b(4):")
	b(4)

	fmt.Println()
	g := 0
	k := &g
	n := c(k)
	m := d()
	fmt.Println(n, m)
	fmt.Println(*k)

	r := e(0)
	r1 := e1(0)
	fmt.Println("R1 := e1(0): ", e(0))
	t := f()
	fmt.Println(r, r1, t)

	i := g1()
	j := h1()
	fmt.Println(i, j)

	bb := b1(4)
	fmt.Println(bb)

	url := "https://www.baidu.com"
	local, l, err := fetch(url)
	fmt.Println(local, l, err)
}

// 无名函数， 先执行return，然后再执行defer
func f1() int {
	var i int
	defer func() {
		i++
		fmt.Println("f11:", i)
	}() // 后面跟着的圆括号相当于匿名函数的参数表，func(){}类似于函数名a，func(){}()就相当于a(),表示函数调用

	defer func() {
		i++
		fmt.Println("f12:", i)
	}()

	i = 1000
	// return i时会先创建一个临时变量，将i的值保存在临时变量中，然后再返回临时变量的值
	// 返回之后再执行defer，但是defer改变的是i的值（i的地址始终不变）,所以不影响返回值
	return i
}

// 有名函数， 先执行return，然后再执行defer
func f2() (i int) {
	defer func() {
		i++
		fmt.Println("f21:", i)
	}()

	defer func() {
		i++
		fmt.Println("f22:", i)
	}()
	i = 1000
	// 有名函数，return i时直接返回i的值，然后执行defer语句，由于i的地址不变，defer语句中改变i的值最终会影响返回值参数i的值，从而影响返回值
	return i
}

// func title(url string) error {
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		return err
// 	}
// 	ct := resp.Header.Get("Content-Type")
// 	if ct != "text/html" && !strings.HasPrefix(ct, "text/html;") {
// 		resp.Body.Close()
// 		return fmt.Errorf("%s has type %s, not text/html", url, ct)
// 	}

// 	doc, err := html.Parse(resp.Body)
// 	resp.Body.Close()
// 	if err != nil {
// 		return fmt.Errorf("parsing %s as HTML: %v", url, err)
// 	}

// 	visitNode := func(n *html.Node) {
// 		if n.Type == html.ElementNode && n.Data == "title" &&
// 			n.FirstChild != nil {
// 			fmt.Println(n.FirstChild.Data)
// 		}
// 	}
// 	// forEachNode(doc, visitNode, nil)
// 	return nil
// }

// func ReadFile(filename string) ([]byte, error) {
// 	f, err := os.Open(filename)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer f.Close()
// 	return ReadAll(f)
// }

var mu sync.Mutex
var m = make(map[string]int)

func lookup(key string) int {
	mu.Lock()
	defer mu.Unlock()
	return m[key]
}

func bigSlowOperation() {
	// 先执行trace()
	// 然后return
	// 再执行trace()()
	defer trace("bigSlowOperation")()

	time.Sleep(1 * time.Second)
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() { log.Printf("exit %s (%s)", msg, time.Since(start)) }
}

func b(x int) (result int) {
	defer func() {
		fmt.Printf("double(%d) = %d", x, result)
	}()
	return x + x
}

func c(x *int) (result int) {
	defer func() {
		(*x)++
	}()
	return *x
}

// 有名函数，但是是对x进行操作，而不是对result进行操作
// 有名函数：返回值参数定义了变量名
func d() (result int) {
	var x int
	defer func() {
		x++
	}()
	return x
}

func e(x int) int {
	defer func() {
		x++
	}()
	return x
}

func e1(x int) int {
	defer func() {
		x++
		fmt.Printf("e11 x:%d\n", x)
	}()

	defer func() {
		x++
		fmt.Printf("e1 x:%d\n", x)
	}()
	x = 0
	return x
}

func f() int {
	var x int
	defer func() {
		x++
	}()
	return x
}

func g1() (result int) {
	defer func() {
		result++
	}()
	return result
}

func h1() int {
	var result int
	defer func() {
		result++
	}()
	return result
}

// a(4)返回值：9
func a(x int) (result int) {
	defer func() {
		result++
		// 输出：double(4)=9
		fmt.Printf("double(%d)=%d\n", x, result)
	}()
	return x + x
}

// 先执行a(4), result==9
// 再执行result = result + x, 此时result == 13
// 再执行赋值语句，bb = b1(4), 最后bb == 13
func b1(x int) (result int) {
	defer func() {
		result += x
	}()
	return a(x)
}

// func c1(filenames []string) {
// 	for _, filename := range filenames {
// 		if err := doFile(filename); err != nil {
// 			return err
// 		}
// 	}
// }

// func doFile(filename string) error {
// 	f, err := os.Open()
// 	if err != nil {
// 		return err
// 	}
// 	defer f.Close(0)
// }

func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	n, err = io.Copy(f, resp.Body)
	if closeErr := f.Close(); err == nil {
		err = closeErr
	}
	return local, n, err
}
