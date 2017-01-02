package main

import (
	"os"

	"github.com/nllptr/pdf"
)

func main() {
	doc := pdf.NewDocument()
	doc.NewLitString("This is some random text.")
	doc.NewLitString("Here is some more text.")
	doc.Write(os.Stdout)
}
