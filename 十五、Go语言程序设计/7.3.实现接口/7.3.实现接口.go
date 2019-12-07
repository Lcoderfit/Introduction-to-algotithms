package main

import (
	"bytes"
	"fmt"
	"io"
	"time"
)

func main() {
	var w io.Writer
	w = os.Stdout
	w = new(bytes.Buffer)

	var rwc io.ReadWriteCloser
	rwc = os.Stdout
	// 编译错误，缺少Close方法：rwc = new(bytes.Buffer)

}

type IntSet struct{}

func (*IntSet) String() string

// var _ = IntSet{}.String()：编译错误

var s IntSet
var _ = s.String()
var _ fmt.Stringer = &s // OK
var _ fmt.Stringer = s  // 编译错误：IntSet缺少String方法

var any interface{}
any = true
any = 12.34
any = "hello"
any = map[string]int{"one": 1}
any = new(bytes.Buffer)

type Artiface interface {
	Title() string
	Creators() []string
	Created() time.Time 
}

type Text interface {
	Pages() int
	Words() int
	PageSize() int
}

type Audio interface {
	Stream() (io.ReadCloser, error)
	RunningTime() time.Duration
	Format() string 
}

type Video interface {
	Stream() (io.ReadCloser, error)
	RunningTime() time.Duration
	Format() string 
	Resolution() (x, y int)
}

type Streamer interface {
	Stream() (io.ReadCloser, error)
	RunningTime() time.Duration
	Format() string 
}

