package pdf

import (
	"fmt"
	"io"
	"sort"
	"strings"
)

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
	/*
		iterera genom strängen
			om balance < 0
				markera indexet som obalanserat
				återställ balansen
		om balansen är positiv n
			ta de n sista öppningsparenteserna och markera dem som obalanserade
		sortera listan med obalanserade index
		inserta backslash i strängen bakifrån
	*/
	balance := 0
	unbalanced := []int{}
	for i, c := range s {
		if c == '(' {
			balance++
		} else if c == ')' {
			balance--
		}
		if balance < 0 {
			unbalanced = append(unbalanced, i)
			balance = 0
		}
	}
	if balance > 0 {
		for i := 0; i < balance; i++ {
			j := len(s)
			for {
				j = strings.LastIndex(s[:j], "(")
				if len(unbalanced) == 0 || unbalanced[len(unbalanced)-1] != j {
					unbalanced = append(unbalanced, j)
					fmt.Println("adding", j)
					break
				}
			}
		}
	}
	sort.Ints(unbalanced)
	for i := len(unbalanced) - 1; i >= 0; i-- {
		s = s[:unbalanced[i]] + "\\" + s[unbalanced[i]:]
	}
	return s
}
