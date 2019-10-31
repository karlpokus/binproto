package bin

import (
	"encoding/binary"
	"io"
	"time"
)

// hasTerm determines if the byte slice contains a line terminator
func hasTerm(b []byte) bool {
	for _, x := range b {
		if x == 10 {
			return true
		}
	}
	return false
}

// readPackets returns Packets read from r with the last one having
// FIN set to true
func readPackets(r io.Reader) (Packets, error) {
	var pkts Packets
	for {
		var b [15]byte
		n, err := r.Read(b[:])
		if err != nil && err != io.EOF {
			return pkts, err
		}
		debugLog.Printf("%d bytes read\n", n)
		isLast := hasTerm(b[:])
		p := Packet{
			Ts:  uint32(time.Now().Unix()),
			Msg: b,
			Fin: isLast,
		}
		pkts = append(pkts, p)
		if isLast {
			break
		}
	}
	return pkts, nil
}

// Encode encodes Packets read from r and writes them to w - one by one
func Encode(w io.Writer, r io.Reader) error {
	pkts, err := readPackets(r)
	if err != nil {
		return err
	}
	debugLog.Printf("%d packets encoded\n", len(pkts))
	return binary.Write(w, binary.BigEndian, pkts) // one by one
}

// Decode decodes-, and returns Packets read from r
func Decode(r io.Reader) (pkts Packets, err error) {
	for {
		var p Packet
		err = binary.Read(r, binary.BigEndian, &p)
		if err != nil {
			return
		}
		pkts = append(pkts, p)
		if p.Fin {
			break
		}
	}
	debugLog.Printf("%d packets decoded\n", len(pkts))
	return
}
