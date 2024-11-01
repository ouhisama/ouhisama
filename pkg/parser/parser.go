package parser

import (
	"log"

	"github.com/ouhisama/ouhisama/pkg/ast"
	"github.com/ouhisama/ouhisama/pkg/token"
)

type parser struct {
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

func (p *parser) want(kind token.TokenKind) token.Token {
	if p.at().Kind != kind {
		log.Fatalf("ERROR Expected the token kind `%v`, but got `%v` instead\n", kind.String(), p.at().Kind.String())
	}
	return p.eat()
}

func (p *parser) isEOF() bool {
	return p.at().Kind == token.EOF
}

func newParser(tokens []token.Token) *parser {
	newTokenLookupTables()
	return &parser{
		tokens:   tokens,
		position: 0,
	}
}

func Parse(tokens []token.Token) ast.BlockStatement {
	p := newParser(tokens)
	body := []ast.Statement{}

	for !p.isEOF() {
		body = append(body, parseStatement(p))
	}

	return ast.BlockStatement{
		Body: body,
	}
}
