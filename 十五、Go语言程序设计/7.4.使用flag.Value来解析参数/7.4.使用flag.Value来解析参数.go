package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	var period = flag.Duration("period", 1*time.Second, "sleep period")

	flag.Parse()
	fmt.Printf("Sleeping for %v...", *period)
	time.Sleep(*period)
	fmt.Println()

	fmt.Println("Begin:")
	time.Sleep(5 * time.Second)
	fmt.Println("Lcoderfit")
}
