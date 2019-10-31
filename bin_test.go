package bin

import (
	"bytes"
	"strings"
	"testing"
)

func TestEncodeDecode(t *testing.T) {
	input := []string{
		"short\n",
		"longer than the maximum\n",
		"really long message that just goes on and on and noone will pay attention ever..\n",
	}
	for _, s := range input {
		r := bytes.NewReader([]byte(s))
		wr := new(bytes.Buffer)
		err := Encode(wr, r)
		if err != nil {
			t.Errorf("expected nil err, got %s", err)
		}
		pkts, err := Decode(wr)
		if err != nil {
			t.Errorf("expected nil err, got %s", err)
		}
		if !strings.Contains(pkts.String(), s) { // because padding in Msg
			t.Errorf("expected %s, got %s", s, pkts)
		}
	}
}
