package pdf

import "testing"

var newLitStringCases = []struct {
	text string
	want LitString
}{
	{
		"this is string 1",
		LitString{Object{1, 0, "this is string 1"}},
	},
	{
		"this is another one (number 2)",
		LitString{Object{2, 0, "this is another one (number 2)"}},
	},
}

func TestNewLitString(t *testing.T) {
	doc := NewDocument()
	for i, c := range newLitStringCases {
		got := doc.NewLitString(c.text)
		if got.num != c.want.num {
			t.Fatalf("Case %d: Wanted num %d, got %d", i+1, c.want.num, got.num)
		}
		if got.gen != 0 {
			t.Fatalf("Case %d: Wanted gen %d, got %d", i+1, c.want.gen, got.gen)
		}
		if got.s != c.want.s {
			t.Fatalf("Case %d: Wanted s '%v', got '%v'", i+1, c.want.s, got.s)
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
