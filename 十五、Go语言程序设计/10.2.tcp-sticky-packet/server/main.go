package main

import (
	"bufio"
	"net"
	""
)

func process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		msg, err := proto.Decode(reader)
	}
}

func main() {

}