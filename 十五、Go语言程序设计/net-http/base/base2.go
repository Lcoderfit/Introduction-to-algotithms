package base

import (
	"fmt"
	"log"
	"net/http"
)

type Engine struct {}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	case "/hello":
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q]=%q\n", k, v)
		}
	default:
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}

func RunBase2() {
	engine := new(Engine)
	//http.ListenAndServe的第二个参数，可以传入任何实现了ServeHTTP接口的实例
	log.Fatal(http.ListenAndServe(":9999", engine))
}