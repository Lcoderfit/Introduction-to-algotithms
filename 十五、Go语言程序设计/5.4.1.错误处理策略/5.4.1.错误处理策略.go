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
	if err := WaitForServer(url); err != nil {
		// log.SetPrefix("wait:")
		// log.SetFlags(0)
		log.Fatalf("Site is down: %v\n", err)
	}
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

func a() {
	if err := Ping(); err != nil {
		log.Printf("ping failed: %v; networking disabled", err)
	}
}

func b() {
	if err := Ping(); err != nil {
		fmt.Fprintf(os.Stderr, "ping failed: %v; networking disabled\n", err)
	}
	dir, err := ioutil.TempDir("", "scratch")
	if err != nil {
		return fmt.Errorf("failed to create temp dir: %v", err)
	}
	os.RemoveAll(dir)
}
