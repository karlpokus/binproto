package bin

import (
	"bytes"
	"testing"
	"strings"
)

func TestEncodeDecode(t *testing.T) {
	data := []string{
		"short",
		"longer than the maximum",
		"really long message that noone will read this side of the year",
	}
	for _, s := range data {
		r := bytes.NewReader([]byte(s))
		wr := new(bytes.Buffer)
		err := Encode(wr, r)
		if err != nil {
			t.Errorf("expected nil err, got %s", err)
		}
		n := len(s)
		i := 0
		j := 15
		if j > n {
			j = n
		}
		for {
			p, err := Decode(wr)
			if err != nil {
				t.Errorf("expected nil err, got %s", err)
			}
			if !strings.Contains(p.String(), s[i:j]) {
				t.Errorf("expected %s, got %s", s[i:j], p)
			}
			if p.IsFin() {
				break
			}
			i = j
			j += 15
			if j > n {
				j = n
			}
		}
	}
}
