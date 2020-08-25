package day3

import (
	"log"
	"net/http"
)

type server int

func (h *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)
	w.Write([]byte("hello world!"))
}

func RunBase() {
	var s server
	http.ListenAndServe(":9999", &s)
}