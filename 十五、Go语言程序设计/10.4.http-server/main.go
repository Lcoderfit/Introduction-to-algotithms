package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func f1(w http.ResponseWriter, r *http.Request) {
	bytes, err := ioutil.ReadFile(`E:\SocialProject\Algorithms-Tags\Introduction-to-algotithms\十五、Go语言程序设计\10.4.http-server\content.txt`)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("read bytes error:%v", err)))
		return
	}
	_, err = w.Write(bytes)
	if err != nil {
		fmt.Println("write data error:", err)
		return
	}
}

func main() {
	// 注册一个处理器
	http.HandleFunc("/hello", f1)
	err := http.ListenAndServe("127.0.0.1:30000", nil)
	if err != nil {
		fmt.Println("server启动失败, error:", err)
		return
	}
}
