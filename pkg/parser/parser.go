package parser

import (
	"fmt"
	"os"

	"github.com/ouhisama/ouhisama/pkg/ast"
	"github.com/ouhisama/ouhisama/pkg/logger"
	"github.com/ouhisama/ouhisama/pkg/token"
)

type parser struct {
	file     string
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
		err, code := logger.NotExpectedTokenError(p.file, p.at(), kind)
		fmt.Println(err)
		os.Exit(int(code))
	}
	return p.eat()
}

func (p *parser) isEOF() bool {
	return p.at().Kind == token.EOF
}

func newParser(file string, tokens []token.Token) *parser {
	newTokenLookupTables()
	return &parser{
		file:     file,
		tokens:   tokens,
		position: 0,
	}
}

func Parse(file string, tokens []token.Token) ast.BlockStatement {
	p := newParser(file, tokens)
	body := []ast.Statement{}

	for !p.isEOF() {
		body = append(body, parseStatement(p))
	}

	return ast.BlockStatement{
		Body: body,
	}
}
