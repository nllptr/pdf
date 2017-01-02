package pdf

import (
	"io"
	"sort"
)

import "strconv"

// NewLitString creates a new literal string.
// TODO: Add some stuff about how octal values can be used to print
// characters outside 7 bit ASCII.
func (d *Document) NewLitString(s string) {
	ls := LitString{
		Object{
			len(d.body) + 1,
			0,
			0,
			s,
		},
	}
	ls.s = strconv.Itoa(ls.num) + " " + strconv.Itoa(ls.gen) + " obj\n"
	ls.s += "(" + balance(s) + ")\nendobj"
	d.body = append(d.body, ls.Object)
}

func (o *Object) write(w io.Writer, offset *int) {
	o.offset = *offset
	bytes, _ := w.Write([]byte(o.s))
	*offset += bytes
}

//func (ls LitString) write(w io.Writer, offset *int) {
/*
	ls.offset = *offset
	bytes, _ := w.Write([]byte(strconv.Itoa(ls.num) + " " + strconv.Itoa(ls.gen) + " obj\n"))
	*offset += bytes
	bytes, _ = w.Write([]byte("(" + balance(ls.s) + ")\nendobj"))
	*offset += bytes
*/
//}

func balance(s string) string {
	stack := []int{}
	backslashList := []int{}

	for i, r := range s {
		if r == '(' {
			stack = append(stack, i)
		} else if r == ')' {
			if len(stack) == 0 {
				backslashList = append(backslashList, i)
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}
	if len(stack) > 0 {
		backslashList = append(backslashList, stack...)
	}
	sort.Ints(backslashList)
	for i := len(backslashList) - 1; i >= 0; i-- {
		s = s[:backslashList[i]] + "\\" + s[backslashList[i]:]
	}
	return s
}
