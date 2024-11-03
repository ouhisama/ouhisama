package parser

import (
	"fmt"
	"strings"

	"github.com/alecthomas/colour"
)

type errorKind uint

const (
	notExpectedToken errorKind = iota + 1
	unexpectedToken
	cannotParseFloat
	noNudHandler
	noLedHandler
)

func (p *parser) error(kind errorKind, msg string, advice string, detail string) {
	file, source := p.file, p.source

	t := p.at()
	index, column, line := t.Position.Value()

	unknown := string(source[index])

	if unknown == "\n" {
		unknown = " "
	}

	rest := ""
	for i, char := range source[index:] {
		if string(char) == "\n" {
			if i == 0 {
				break
			} else {
				rest = source[index+uint(len(unknown)):index+uint(i)]
				break
			}
		} else if i == len(source[index:])-1 {
			rest = source[index+uint(len(unknown)):index + uint(i+1)]
			break
		}
	}

	err := fmt.Sprintf("P%0*d", 3, kind)
	code := colour.Sprintf("^7%v^1%v^7%v", source[index+1-column:index], unknown, rest)

	hint := colour.Sprintf("%v^1%v %v^R", strings.Repeat(" ", int(column)-1), strings.Repeat("^", int(t.Length)), advice)

	if detail != "" {
		detail = "\n" + detail + "\n"
	}

	fmt.Println(colour.Sprintf("^1ERROR[%v]^7 %v\n\n\t> %v\n\t|\n%v\t| %v\n\t| %v\n%v", err, msg, file, line, code, hint, detail))
}
