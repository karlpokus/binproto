package main

import (
	"fmt"
	"net"
	"time"

	"bin"
)

func handler(conn net.Conn) {
	defer conn.Close()
	for {
		conn.SetDeadline(time.Now().Add(30 * time.Second))
		p, err := bin.Decode(conn)
		if err != nil {
			fmt.Errorf("Decode err:%s\n", err)
			continue
		}
		// leaving off the line terminator in fmt will print the entire client msg
		// on the same line even beyond the 12 bytes until recieving a line terminator
		fmt.Printf("%s", p)
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
