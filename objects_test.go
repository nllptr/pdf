package pdf

import "testing"

import "bytes"

var newLitStringCases = []struct {
	text string
	want LitString
}{
	{
		"this is string 1",
		LitString{Object{1, 0, 0, "1 0 obj\n(this is string 1)\nendobj"}},
	},
	{
		"this is another one (number 2)",
		LitString{Object{2, 0, 0, "2 0 obj\n(this is another one (number 2))\nendobj"}},
	},
}

func TestNewLitString(t *testing.T) {
	doc := NewDocument()
	for i, c := range newLitStringCases {
		doc.NewLitString(c.text)
		if doc.body[i].num != c.want.num {
			t.Fatalf("Case %d: Wanted num %d, got %d", i+1, c.want.num, doc.body[i].num)
		}
		if doc.body[i].gen != 0 {
			t.Fatalf("Case %d: Wanted gen %d, got %d", i+1, c.want.gen, doc.body[i].gen)
		}
		if doc.body[i].s != c.want.s {
			t.Fatalf("Case %d: Wanted s '%v', got '%v'", i+1, c.want.s, doc.body[i].s)
		}
	}
}

var balanceCases = []struct {
	input string
	want  string
}{
	{"((hello))", "((hello))"},
	{"((hello)", "\\((hello)"},
	{"(hello))", "(hello)\\)"},
	{"(there(is)no()spoon)()(", "(there(is)no()spoon)()\\("},
	{"(fix(this(", "\\(fix\\(this\\("},
	{") odd stuff ()", "\\) odd stuff ()"},
	{")(", "\\)\\("},
}

func TestBalance(t *testing.T) {
	for i, c := range balanceCases {
		got := balance(c.input)
		if got != c.want {
			t.Fatalf("Case %d: Wanted '%v', got '%v'", i+1, c.want, got)
		}
	}
}

var WriteLitStrCases = []struct {
	input string
	want  string
}{
	{
		"This is (just) a test :)",
		"1 0 obj\n(This is (just) a test :\\))\nendobj",
	},
}

func TestWriteLitString(t *testing.T) {
	doc := NewDocument()
	for i, c := range WriteLitStrCases {
		doc.NewLitString(c.input)
		var buf bytes.Buffer
		dummy := 0
		doc.body[i].write(&buf, &dummy)
		if buf.String() != c.want {
			t.Fatalf("Case %d:\nWanted\n%v\n\nGot\n%v", i, c.want, buf.String())
		}
	}
}
