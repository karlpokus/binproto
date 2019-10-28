package bin

import (
	"encoding/binary"
	"fmt"
	"io"
	"time"
)

type packet struct { // 16 bytes
	Ts  uint32
	Msg [12]byte
}

func (p packet) String() string {
	return fmt.Sprintf("%s", p.Msg)
}

func (p packet) Time() time.Time {
	return time.Unix(int64(p.Ts), 0)
}

// Encode encodes a packet read from r and writes to w
func Encode(w io.Writer, r io.Reader) error {
	var b [12]byte
	_, err := r.Read(b[:])
	if err != nil {
		return err
	}
	p := packet{
		Ts:  uint32(time.Now().Unix()),
		Msg: b,
	}
	return binary.Write(w, binary.BigEndian, &p)
}

// Decode decodes a packet read from r
func Decode(r io.Reader) (packet, error) {
	var p packet
	err := binary.Read(r, binary.BigEndian, &p)
	return p, err
}
