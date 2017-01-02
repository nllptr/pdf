package pdf

import "io"

type Document struct {
	header  Header
	body    Body
	xref    Xref
	trailer Trailer
	cOffset int
}

type Header struct {
	version        string
	containsBinary bool // TODO: If the file contains binary data, the first line should be followed by a line of at least 4 binary (>128) characters
}

// Body contains Objects or Object streams, which are sequences of objects.
type Body []Object

type Xref int

type Trailer string

// Object represents indirect objects according to the PDF spec. An
// indirect object has an object number and a generation number.
type Object struct {
	num    int
	gen    int
	offset int
	s      string
}

// ObjectWriter declares the write method that writes the object.
type ObjectWriter interface {
	write(io.Writer, *int)
}

// LitString represents the "Literal string" according to the PDF spec.
type LitString struct {
	Object
}
