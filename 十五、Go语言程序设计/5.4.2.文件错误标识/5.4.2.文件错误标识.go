package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

func main() {
	var EOF = errors.New("EOF")
	fmt.Println(io.EOF, EOF)
	in := bufio.NewReader(os.Stdin)
	for {
		r, p, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Errorf("read failed: %v", err)
		}
		fmt.Println("r&P::", r, p)
	}
}
