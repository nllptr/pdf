package pdf

import (
	"io"
	"strconv"
)

// NewDocument creates a new document with default values.
func NewDocument() Document {
	doc := Document{}
	doc.header.version = "1.7"
	return doc
}

func (d *Document) Write(w io.Writer) error {
	d.header.write(w)
	d.body.write(w)

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

func (b *Body) write(w io.Writer) {
	for _, obj := range *b {
		obj.write(w)
		w.Write([]byte("\n\n"))
	}
}

func (d *Document) generateXref() string {
	/*
		xref
		förstaobjnr antalobj
		0000000000 65535 f <<-- head of free objects linked list
		nnnnnnnnnn ggggg n eol <<-- för varje objekt

		nnnnnnnnnn 10 siffror offset från början av filen till objekt
		ggggg generationsnumer, alltid 0 vid skapande
		n/f n=in use, f=free (vid modifiering, ej aktuellt)
		eol CR+newline eller space + en eol character " \n"
	*/
	xref := "xref\n" +
		"0 " + strconv.Itoa(len(d.body)) + "\n" +
		"0000000000 65535 f\n"
	return xref
}
