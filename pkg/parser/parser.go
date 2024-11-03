package parser

import (
	"fmt"
	"os"

	"github.com/ouhisama/ouhisama/pkg/ast"
	"github.com/ouhisama/ouhisama/pkg/token"
)

type parser struct {
	file     string
	source   string
	tokens   []token.Token
	position uint
}

func (p *parser) at() token.Token {
	return p.tokens[p.position]
}

func (p *parser) eat() token.Token {
	t := p.at()
	p.position++
	return t
}

func (p *parser) next() token.Token {
	return p.tokens[p.position+1]
}

func (p *parser) advance() {
	p.position++
}

func (p *parser) previous() token.Token {
	if p.position != 0 {
		return p.tokens[p.position-1]
	} else {
		// Program must't get here
		os.Exit(1)
		return token.Token{}
	}
}

func (p *parser) want(kind token.TokenKind) token.Token {
	if p.at().Kind != kind {
		switch kind {
		case token.Newline:
			msg := "Unexpected newline"
			advice := "remove the newline"
			p.error(notExpectedToken, msg, advice, "")
			os.Exit(1)
		default:
			msg := fmt.Sprintf("Expected a `%v`, instead of `%v`", kind.String(), p.at().Kind.String())
			advice := fmt.Sprintf("this's supposed to be a `%v`", kind.String())
			p.error(notExpectedToken, msg, advice, "")
			os.Exit(1)
		}
	}
	return p.eat()
}

func (p *parser) isEOF() bool {
	return p.at().Kind == token.EOF
}

func newParser(file string, source string, tokens []token.Token) *parser {
	newTokenLookupTables()
	return &parser{
		file:     file,
		source:   source,
		tokens:   tokens,
		position: 0,
	}
}

func Parse(file string, source string, tokens []token.Token) ast.BlockStatement {
	p := newParser(file, source, tokens)
	body := []ast.Statement{}

	for !p.isEOF() {
		body = append(body, parseStatement(p))
	}

	return ast.BlockStatement{
		Body: body,
	}
}
