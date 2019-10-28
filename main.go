package main

import (
	"fmt"
	"encoding/binary"
	"time"
	"bytes"
)

type packet struct {
  Ts uint32
	Msg [12]byte
}

func encode(s string) (*bytes.Buffer, error) {
	p := packet{
  	Ts: uint32(time.Now().Unix()),
  }
	copy(p.Msg[:], s)
  buf := new(bytes.Buffer) // var bytes.Buffer does not implement io.Writer
  err := binary.Write(buf, binary.BigEndian, &p)
	return buf, err
}

func decode(buf *bytes.Buffer) (packet, error) {
	var p packet
  err := binary.Read(buf, binary.BigEndian, &p)
  return p, err
}

func main() {
	buf, err := encode("hi i am bixa and I am a kitty")
	if err != nil {
		panic(err)
	}
	p, err := decode(buf)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v %s\n", time.Unix(int64(p.Ts), 0), p.Msg)
}
