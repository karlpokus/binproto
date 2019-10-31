package bin

import (
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

func StartClient(debug bool) error {
	debugMode(debug)
	conn, err := net.DialTimeout("tcp", ":19200", time.Duration(3e9))
	if err != nil {
		return err
	}
	fmt.Println("client connected")
	if debug {
		go func() {
			io.Copy(os.Stdout, conn)
		}()
	}
	for {
		err = Encode(conn, os.Stdin)
		if err != nil {
			fmt.Printf("Encoding error: %s\n", err)
			continue
		}
	}
	return nil
}
