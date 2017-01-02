package pdf

import (
	"fmt"
	"io"
	"log"
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
	w.Write([]byte("%%EOF"))
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
		//w.Write(binaryBytes)
		//w.Write([]byte("\n"))
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
	log.Println(d.cOffset)
	xref := "xref\n" +
		"0 " + strconv.Itoa(len(d.body)+1) + "\n" +
		"0000000000 65535 f\n"
	w.Write([]byte(xref))
	for _, o := range d.body {
		fmt.Fprintf(w, "%010d 00000 n \n", o.offset)
		//xref += strconv.Itoa(o.offset)
		//xref += " 00000 n \n"
	}
	//w.Write([]byte(xref))
}
