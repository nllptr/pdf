package pdf

import (
	"fmt"
	"io"
	"strconv"
)

// NewDocument creates a new document with default values.
func NewDocument() Document {
	doc := Document{}
	doc.header.version = "1.7"
	doc.cOffset = 0
	return doc
}

func (d *Document) Write(w io.Writer) error {
	d.cOffset += d.header.write(w)
	d.body.write(w, &d.cOffset)
	d.writeXref(w)
	d.writeTrailer(w)
	return nil
}

func (h *Header) write(w io.Writer) int {
	offset := 0
	bytes, _ := w.Write([]byte("%PDF-" + h.version + "\n"))
	offset += bytes
	if h.containsBinary {
		binaryBytes := make([]byte, 4)
		for i := range binaryBytes {
			binaryBytes[i] = 128
		}
		bytes, _ = w.Write(append([]byte("%"), append(binaryBytes, '\n')...))
		offset += bytes
	}
	return offset
}

func (b *Body) write(w io.Writer, offset *int) {
	for i, o := range *b {
		o.write(w, offset)
		(*b)[i] = o
		bytes, _ := w.Write([]byte("\n\n"))
		*offset += bytes
	}
}

func (d *Document) writeXref(w io.Writer) {
	d.xref = Xref(d.cOffset)
	xref := "xref\n" +
		"0 " + strconv.Itoa(len(d.body)+1) + "\n" +
		"0000000000 65535 f\n"
	bytes, _ := w.Write([]byte(xref))
	d.cOffset += bytes
	for _, o := range d.body {
		bytes, _ = fmt.Fprintf(w, "%010d 00000 n \n", o.offset)
		d.cOffset += bytes
	}

	bytes, _ = w.Write([]byte("\n"))
	d.cOffset += bytes
}

func (d *Document) writeTrailer(w io.Writer) {
	w.Write([]byte("trailer\n"))
	/*
		TODO: Trailer doctionary
		Required: "Size"; equals highest object number + 1
				  "Root"; Catalog dictionary
				  "ID"; Strongly recommended
	*/
	w.Write([]byte("<<\n"))
	w.Write([]byte("/Size " + strconv.Itoa(len(d.body)+1) + "\n"))
	w.Write([]byte(">>\n"))
	w.Write([]byte("startxref\n"))
	fmt.Fprintf(w, "%d\n", d.xref)
	w.Write([]byte("%%EOF"))
}
