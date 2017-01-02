package pdf

import (
	"bytes"
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	doc := NewDocument()
	fmt.Println("version", doc.header.version)
}

var headerWriteCases = []struct {
	h         Header
	want      string
	wantBytes int
}{
	{
		Header{
			"1.7",
			false,
		},
		"%PDF-1.7\n",
		9,
	},
	{
		Header{
			"1.7",
			true,
		},
		"%PDF-1.7\n%" + string([]byte{128, 128, 128, 128}) + "\n",
		15,
	},
}

func TestHeaderWrite(t *testing.T) {
	for i, c := range headerWriteCases {
		var b bytes.Buffer
		gotBytes := c.h.write(&b)
		got := b.String()
		if got != c.want {
			t.Fatalf("Case %d:\nWanted:\n%v\n\nGot:\n%v\n\n", i+1, c.want, got)
		}
		if gotBytes != c.wantBytes {
			t.Fatalf("Case %d: Wanted %d bytes, got %d bytes", i+1, c.wantBytes, gotBytes)
		}
	}
}
