package tokeniser

import (
	"fmt"
	"os"
	"regexp"

	"github.com/alecthomas/colour"
	"github.com/ouhisama/ouhisama/pkg/token"
)

type handler func(tokeniser *tokeniser, regex *regexp.Regexp)

type pattern struct {
	regex   *regexp.Regexp
	handler handler
}

type position struct {
	index  uint
	column uint
	line   uint
}

func (p position) value() (uint, uint, uint) {
	return p.index, p.column, p.line
}

type tokeniser struct {
	Tokens   []token.Token
	patterns []pattern
	file     string
	source   string
	position position
}

func (t *tokeniser) push(token token.Token) {
	t.Tokens = append(t.Tokens, token)
}

func (t *tokeniser) remainder() string {
	return t.source[t.position.index:]
}

func (t *tokeniser) isEOF() bool {
	return t.position.index >= uint(len(t.source))
}

func (t *tokeniser) advanceIndex(length uint) {
	t.position.index += length
}

func (t *tokeniser) advanceLine(length uint) {
	t.position.line += length
	t.position.column = 1
	t.advanceIndex(length)
}

func (t *tokeniser) advanceColumn(length uint) {
	t.position.column += length
	t.advanceIndex(length)
}

func defaultHandler(kind token.TokenKind, value string) handler {
	return func(t *tokeniser, regex *regexp.Regexp) {
		index, column, line := t.position.value()
		t.advanceColumn(uint(len(value)))
		t.push(token.NewToken(kind, value, token.NewTokenPosition(index, column, line), uint(len(value))))
	}
}

func newlineHandler(t *tokeniser, regex *regexp.Regexp) {
	matched := regex.FindString(t.remainder())
	index, column, line := t.position.value()
	t.push(token.NewToken(token.Newline, token.Newline.String(), token.NewTokenPosition(index, column, line), uint(len(matched))))
	t.advanceLine(uint(len(matched)))
}

func indentationHandler(t *tokeniser, regex *regexp.Regexp) {
	matched := regex.FindString(t.remainder())
	index, column, line := t.position.value()
	t.push(token.NewToken(token.Indentation, token.Indentation.String(), token.NewTokenPosition(index, column, line), uint(len(matched))))
	t.advanceColumn(uint(len(matched)))
}

func whitespaceHandler(t *tokeniser, regex *regexp.Regexp) {
	matched := regex.FindString(t.remainder())
	index, column, line := t.position.value()
	t.push(token.NewToken(token.Whitespace, token.Whitespace.String(), token.NewTokenPosition(index, column, line), uint(len(matched))))
	t.advanceColumn(uint(len(matched)))
}

func numberHandler(t *tokeniser, regex *regexp.Regexp) {
	matched := regex.FindString(t.remainder())
	index, column, line := t.position.value()
	t.push(token.NewToken(token.Number, string(matched), token.NewTokenPosition(index, column, line), uint(len(matched))))
	t.advanceColumn(uint(len(matched)))
}

func newTokeniser(file string, source string) *tokeniser {
	return &tokeniser{
		Tokens: []token.Token{},
		patterns: []pattern{
			{regexp.MustCompile(`\n`), newlineHandler},
			{regexp.MustCompile(`\t+`), indentationHandler},
			{regexp.MustCompile(`\s+`), whitespaceHandler},

			{regexp.MustCompile(`[0-9]+(\.[0-9]+)?`), numberHandler},

			{regexp.MustCompile(`=`), defaultHandler(token.Equal, "=")},
			{regexp.MustCompile(`\+`), defaultHandler(token.Plus, "+")},
			{regexp.MustCompile(`\-`), defaultHandler(token.Hyphen, "-")},
			{regexp.MustCompile(`\*`), defaultHandler(token.Star, "*")},
			{regexp.MustCompile(`\/`), defaultHandler(token.Slash, "/")},
			{regexp.MustCompile(`\%`), defaultHandler(token.Percent, "%")},
			{regexp.MustCompile(`\#`), defaultHandler(token.Hashtag, "#")},
			{regexp.MustCompile(`\(`), defaultHandler(token.LBracket, "(")},
			{regexp.MustCompile(`\)`), defaultHandler(token.RBracket, ")")},
		},
		file:   file,
		source: source,
		position: position{
			index:  0,
			column: 1,
			line:   1,
		},
	}
}

func Tokenise(file string, source string) []token.Token {
	t := newTokeniser(file, source)

	for !t.isEOF() {
		matched := false

		for _, pattern := range t.patterns {
			location := pattern.regex.FindStringIndex(t.remainder())

			if location != nil && location[0] == 0 {
				pattern.handler(t, pattern.regex)
				matched = true
				break
			}
		}

		if !matched {
			msg := fmt.Sprintf("Got an unrecongnised token `%v` while tokenising", string(t.source[t.position.index]))
			advice := colour.Sprintf("try removing this ^Sstupid^R^1 unrecongnised token")
			t.error(unrecongnisedToken, msg, advice, "")
			os.Exit(1)
		}
	}

	index, column, line := t.position.value()
	t.push(token.NewToken(token.EOF, token.EOF.String(), token.NewTokenPosition(index, column, line), 1))
	return t.Tokens
}
