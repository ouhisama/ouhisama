package tokeniser

import (
	"log"
	"regexp"
	"strconv"
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

type tokeniser struct {
	Tokens   []Token
	patterns []pattern
	source   string
	position position
}

func (t *tokeniser) at(n uint) byte {
	return t.source[t.position.index+n]
}

func (t *tokeniser) push(token Token) {
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

func defaultHandler(kind TokenKind, value TokenValue) handler {
	return func(t *tokeniser, regex *regexp.Regexp) {
		t.advanceColumn(uint(len(value)))
		t.push(newToken(kind, value))
	}
}

func newlineHandler(t *tokeniser, regex *regexp.Regexp) {
	matched := regex.FindString(t.remainder())
	t.push(newToken(Newline, TokenValue(strconv.Itoa(len(matched)))))
	t.advanceLine(uint(len(matched)))
}

func indentationHandler(t *tokeniser, regex *regexp.Regexp) {
	matched := regex.FindString(t.remainder())
	t.push(newToken(Indentation, TokenValue(strconv.Itoa(len(matched)/4))))
	t.advanceColumn(uint(len(matched)))
}

func whitespaceHandler(t *tokeniser, regex *regexp.Regexp) {
	matched := regex.FindString(t.remainder())
	t.push(newToken(Whitespace, TokenValue(strconv.Itoa(len(matched)))))
	t.advanceColumn(uint(len(matched)))
}

func numberHandler(t *tokeniser, regex *regexp.Regexp) {
	matched := regex.FindString(t.remainder())
	t.push(newToken(Number, TokenValue(matched)))
	t.advanceColumn(uint(len(matched)))
}

func newTokeniser(source string) *tokeniser {
	return &tokeniser{
		Tokens: []Token{},
		patterns: []pattern{
			{regexp.MustCompile(`\n`), newlineHandler},
			{regexp.MustCompile(`\t+`), indentationHandler},
			{regexp.MustCompile(`\s+`), whitespaceHandler},

			{regexp.MustCompile(`-?[0-9]+(\.[0-9]+)?`), numberHandler},

			{regexp.MustCompile(`=`), defaultHandler(Equal, "=")},
			{regexp.MustCompile(`\+`), defaultHandler(Plus, "+")},
			{regexp.MustCompile(`\-`), defaultHandler(Hyphen, "-")},
			{regexp.MustCompile(`\*`), defaultHandler(Star, "*")},
			{regexp.MustCompile(`\/`), defaultHandler(Slash, "/")},
			{regexp.MustCompile(`\%`), defaultHandler(Percent, "%")},
			{regexp.MustCompile(`\#`), defaultHandler(Hashtag, "#")},
			{regexp.MustCompile(`\(`), defaultHandler(LBracket, "(")},
			{regexp.MustCompile(`\)`), defaultHandler(RBracket, ")")},
		},
		source: source,
		position: position{
			index:  0,
			column: 1,
			line:   1,
		},
	}
}

func Tokenise(source string) []Token {
	t := newTokeniser(source)

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
			coloumn, line := t.position.column, t.position.line
			character := string(t.source[t.position.index])
			log.Fatalf("ERROR Unrecongnised token `%v` at line %v column %v\n", character, line, coloumn)
		}
	}

	t.push(newToken(EOF, ""))
	return t.Tokens
}
