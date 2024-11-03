package tokeniser

import (
	"fmt"
	"strings"

	"github.com/alecthomas/colour"
	"github.com/ouhisama/ouhisama/pkg/token"
)

type errorKind uint

const (
	unrecongnisedToken errorKind = iota + 1
)

func (t *tokeniser) error(kind errorKind, msg string, advice string, detail string) {
	file, source := t.file, t.source

	index, column, line := t.position.value()

	unknown := string(source[index])

	var end uint
	for i, char := range source[index:] {
		if string(char) == "\n" {
			end = index + uint(i)
			break
		} else if i == len(source[index:])-1 {
			end = index + uint(i+1)
			break
		}
	}

	err := fmt.Sprintf("T%0*d", 3, kind)
	code := colour.Sprintf("^7%v^1%v^7%v", source[index+1-column:index], unknown, source[index+uint(len(unknown)):end])

	var indentations uint
	if t.previous().Kind == token.Indentation {
		indentations = t.previous().Length
	}
	hint := colour.Sprintf("%v^1%v %v^R", strings.Repeat("\t", int(indentations)), strings.Repeat("^", len(unknown)), advice)

	if detail != "" {
		detail = "\n" + detail + "\n"
	}

	fmt.Println(colour.Sprintf("^1ERROR[%v]^7 %v\n\n\t> %v\n\t|\n%v\t| %v\n\t| %v\n%v", err, msg, file, line, code, hint, detail))
}
