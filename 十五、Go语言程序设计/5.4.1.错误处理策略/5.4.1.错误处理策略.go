package main

import (
	"fmt"
	"net/http"
	"time"

	"golang.org/x/net/html"
)

func main() {
	url := "https://www.baidu.com/"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("ERROR:", err)
	} else {
		fmt.Println(resp)
		doc, err := html.Parse(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Errorf("parsing %s as HTML: %v", url, err)
		} else {
			fmt.Println("Success:", doc)
		}
	}
	error := WaitForServer(url)
	fmt.Println(error)
}

func WaitForServer(url string) error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err == nil {
			return nil
		}
		fmt.Printf("server not responding (%s); retrying...", err)
		time.Sleep(time.Second << uint(tries)) // 指数退避策略
	}
	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}
