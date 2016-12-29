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
	h    Header
	want string
}{
	{
		Header{
			"1.7",
			false,
		},
		"%PDF-1.7\n",
	},
	{
		Header{
			"1.7",
			true,
		},
		"%PDF-1.7\n%" + string([]byte{128, 128, 128, 128}) + "\n",
	},
}

func TestHeaderWrite(t *testing.T) {
	for i, c := range headerWriteCases {
		var b bytes.Buffer
		c.h.write(&b)
		got := b.String()
		if got != c.want {
			t.Fatalf("Case %d:\nWanted:\n%v\n\nGot:\n%v\n\n", i+1, c.want, got)
		}
	}
}
