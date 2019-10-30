package bin

import (
	"fmt"
	"time"
)

// Packet fields must be exported
// to work with binary.Read and Write
type Packet struct { // 20 bytes
	Fin bool   // 1 byte
	Ts  uint32 // 4 bytes
	Msg [15]byte
}

func (p Packet) String() string {
	return fmt.Sprintf("%s", p.Msg)
}

func (p Packet) Time() time.Time {
	return time.Unix(int64(p.Ts), 0)
}

type Packets []Packet

func (pkts Packets) String() string {
	var out string
	for _, p := range pkts {
		out += fmt.Sprintf("%s", p.Msg)
	}
	return out
}
