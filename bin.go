package bin

import (
	"encoding/binary"
	"fmt"
	"io"
	"time"
)

type packet struct { // 20 bytes
	Fin bool   // 1 byte
	Ts  uint32 // 4 bytes
	Msg [15]byte
}

func (p packet) String() string {
	return fmt.Sprintf("%s", p.Msg)
}

func (p packet) Time() time.Time {
	return time.Unix(int64(p.Ts), 0)
}

func (p packet) IsFin() bool {
	return p.Fin
}

// splitPackets returns n packets read from r with the last one having
// FIN set to true
func splitPackets(r io.Reader) ([]packet, error) {
	var packets []packet
	for {
		var b [15]byte
		n, err := r.Read(b[:])
		if err != nil && err != io.EOF {
			return packets, err
		}
		if n == 0 {
			break
		}
		p := packet{
			Ts:  uint32(time.Now().Unix()),
			Msg: b,
		}
		packets = append(packets, p)
	}
	packets[len(packets)-1].Fin = true // set FIN true on last packet
	return packets, nil
}

// Encode encodes n packets read from r and writes them to w - one by one
func Encode(w io.Writer, r io.Reader) error {
	packets, err := splitPackets(r)
	if err != nil {
		return err
	}
	return binary.Write(w, binary.BigEndian, packets) // one by one
}

// Decode decodes a packets read from r
func Decode(r io.Reader) (packet, error) {
	var p packet
	err := binary.Read(r, binary.BigEndian, &p)
	return p, err
}
