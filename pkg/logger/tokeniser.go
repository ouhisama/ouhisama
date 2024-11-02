package logger

import (
	"fmt"
	"strings"

	"github.com/alecthomas/colour"
)

func UnrecongnisedTokenError(file string, source string, unknown string, index uint, line uint, column uint) (string, TokeniserErrorCode) {
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

	err := fmt.Sprintf("T%0*d", 3, UnrecongnisedToken)

	code := colour.Sprintf("^7%v^1%v^7%v", source[index+1-column:index], unknown, source[index+uint(len(unknown)):end])

	advice := colour.Sprintf("%v^1%v try removing this ^Sstupid^R^1 unrecongnised token", strings.Repeat(" ", int(column)-1), strings.Repeat("^", len(unknown)))

	return colour.Sprintf("^1ERROR[%v]^7 Got an unrecongnised token `%v` while tokenising\n\n\t> %v\n\t|\n%v\t| %v\n\t| %v\n", err, unknown, file, line, code, advice), UnrecongnisedToken
}
