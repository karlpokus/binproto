package bin

import (
	"bytes"
	"testing"
)

func TestEncodeDecode(t *testing.T) {
	s := "helloy there" // 12 bytes
	r := bytes.NewReader([]byte(s))
	wr := new(bytes.Buffer)
	err := Encode(wr, r)
	if err != nil {
		t.Errorf("expected nil err, got %s", err)
	}
	p, err := Decode(wr)
	if err != nil {
		t.Errorf("expected nil err, got %s", err)
	}
	if p.String() != s {
		t.Errorf("want %s, got %s", s, p)
	}
}
