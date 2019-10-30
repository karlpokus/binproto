package main

import (
	"fmt"
	"net"
	"time"
	"io"

	"bin"
)

func handler(conn net.Conn) {
	fmt.Println("Client connected")
	defer conn.Close()
	for {
		conn.SetDeadline(time.Now().Add(30 * time.Second))
		pkts, err := bin.Decode(conn)
		if err, ok := err.(net.Error); ok && err.Timeout() {
			fmt.Println("Connection timeout")
			return
		}
		if err == io.EOF {
			fmt.Println("Client disconnected")
			return
		}
		if err != nil {
			fmt.Errorf("Decoding error: %s\n", err)
			continue
		}
		fmt.Printf("%s", pkts) // payload includes a line terminator
		fmt.Fprintln(conn, "ok")
	}
}

func start() error {
	l, err := net.Listen("tcp", ":19200")
	if err != nil {
		return err
	}
	fmt.Println("server listening")
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Printf("connection error: %s\n", err)
			continue
		}
		go handler(conn)
	}
}

func main() {
	err := start()
	if err != nil {
		fmt.Printf("%s\n", err)
	}
}
