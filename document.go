package pdf

import (
	"io"
)

// NewDocument creates a new document with default values.
func NewDocument() Document {
	doc := Document{}
	doc.header.version = "1.7"
	return doc
}

func (d *Document) Write(w io.Writer) error {
	d.header.write(w)
	w.Write([]byte("%%EOF"))
	return nil
}

func (h *Header) write(w io.Writer) {
	w.Write([]byte("%PDF-" + h.version + "\n"))
	if h.containsBinary {
		binaryBytes := make([]byte, 4)
		for i := range binaryBytes {
			binaryBytes[i] = 128
		}
		w.Write([]byte("%"))
		w.Write(binaryBytes)
		w.Write([]byte("\n"))
	}
}
