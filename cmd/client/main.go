package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"time"

	"bin"
)

func main() {
	conn, err := net.DialTimeout("tcp", ":19200", time.Duration(3e9))
	if err != nil {
		fmt.Printf("Connection error: %s\n", err)
		return
	}
	fmt.Println("client connected")
	go func() {
		io.Copy(os.Stdout, conn)
	}()
	for {
		err = bin.Encode(conn, os.Stdin)
		if err != nil {
			fmt.Printf("Encoding error: %s\n", err)
			break
		}
	}
}
