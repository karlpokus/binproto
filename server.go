package bin

import (
	"fmt"
	"io"
	"net"
	"time"
)

func StartServer(debug bool) error {
	debugMode(debug)
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

func handler(conn net.Conn) {
	fmt.Println("client connected")
	defer conn.Close()
	for {
		conn.SetDeadline(time.Now().Add(60e9))
		pkts, err := Decode(conn)
		if err, ok := err.(net.Error); ok && err.Timeout() {
			fmt.Println("connection timeout")
			return
		}
		if err == io.EOF {
			fmt.Println("client disconnected")
			return
		}
		if err != nil {
			fmt.Printf("decoding error: %s\n", err)
			continue
		}
		fmt.Printf("%s", pkts) // payload includes line terminator
		fmt.Fprintln(conn, "[ACK]")
	}
}
