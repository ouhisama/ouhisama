package logger

import (
	"fmt"
	"os"
	"strings"

	"github.com/alecthomas/colour"
	"github.com/ouhisama/ouhisama/pkg/token"
)

func UnexpectedTokenError(file string, t token.Token) (string, ParserErrorCode) {
	bytes, _ := os.ReadFile(file)
	source := string(bytes)

	index, column, line := t.Position.Index, t.Position.Column, t.Position.Line

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

	err := fmt.Sprintf("P%0*d", 3, UnexpectedToken)

	code := colour.Sprintf("^7%v^1%v^7%v", source[index+1-column:index], t.Value, source[index+t.Length:end])

	advice := colour.Sprintf("%v^1%v check your code syntax or just remove this", strings.Repeat(" ", int(column)-1), strings.Repeat("^", len(t.Value)))

	return colour.Sprintf("^1ERROR[%v]^7 Got an unexpected token `%v` while parsing an expression\n\n\t> %v\n\t|\n%v\t| %v\n\t| %v\n", err, t.Value, file, line, code, advice), UnexpectedToken
}

func NotExpectedTokenError(file string, got token.Token, want token.TokenKind) (string, ParserErrorCode) {
	bytes, _ := os.ReadFile(file)
	source := string(bytes)

	index, column, line := got.Position.Index, got.Position.Column, got.Position.Line

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

	err := fmt.Sprintf("P%0*d", 3, NotExpectedToken)

	code := colour.Sprintf("^7%v^1%v^7%v", source[index+1-column:index], got.Value, source[index+got.Length:end])

	advice := colour.Sprintf("%v^1%v this's supposed to be a `%v`", strings.Repeat(" ", int(column)-1), strings.Repeat("^", len(got.Value)), want.String())

	return colour.Sprintf("^1ERROR[%v]^7 Expected a `%v`, but got a `%v` instead while parsing\n\n\t> %v\n\t|\n%v\t| %v\n\t| %v\n", err, want.String(), got.Kind.String(), file, line, code, advice), NotExpectedToken
}

func NoNudHandlerError(file string, t token.Token) (string, ParserErrorCode) {
	bytes, _ := os.ReadFile(file)
	source := string(bytes)

	index, column, line := t.Position.Index, t.Position.Column, t.Position.Line

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

	err := fmt.Sprintf("P%0*d", 3, NoNudHandler)

	code := colour.Sprintf("^7%v^1%v^7%v", source[index+1-column:index], t.Value, source[index+t.Length:end])

	advice := colour.Sprintf("%v^1%v you might put an incorrect stuff ^Slike you^R^1 here", strings.Repeat(" ", int(column)-1), strings.Repeat("^", len(t.Value)))

	return colour.Sprintf("^1ERROR[%v]^7 No left denotation handler for the token `%v` while parsing an expression\n\n\t> %v\n\t|\n%v\t| %v\n\t| %v\n", err, t.Value, file, line, code, advice), NoNudHandler
}

func NoLedHandlerError(file string, t token.Token) (string, ParserErrorCode) {
	bytes, _ := os.ReadFile(file)
	source := string(bytes)

	index, column, line := t.Position.Index, t.Position.Column, t.Position.Line

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

	err := fmt.Sprintf("P%0*d", 3, NoLedHandler)

	code := colour.Sprintf("^7%v^1%v^7%v", source[index+1-column:index], t.Value, source[index+t.Length:end])

	advice := colour.Sprintf("%v^1%v you might put an incorrect stuff ^Slike you^R^1 here", strings.Repeat(" ", int(column)-1), strings.Repeat("^", len(t.Value)))

	return colour.Sprintf("^1ERROR[%v]^7 No left denotation handler for the token `%v` while parsing an expression\n\n\t> %v\n\t|\n%v\t| %v\n\t| %v\n", err, t.Value, file, line, code, advice), NoLedHandler
}
