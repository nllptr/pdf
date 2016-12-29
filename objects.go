package pdf

import "io"
import "sort"

func (d *Document) NewLitString(s string) LitString {
	ls := LitString{
		Object{
			len(d.body) + 1,
			0,
			s,
		},
	}
	d.body = append(d.body, ls.Object)
	return ls
}

func (ls LitString) write(w io.Writer) {
	ls.s = balance(ls.s)
	//w.Write()
}

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
