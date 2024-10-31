package parser

import (
	"github.com/ouhisama/ouhisama/pkg/ast"
	"github.com/ouhisama/ouhisama/pkg/token"
)

type bindingPower uint

const (
	zero bindingPower = iota
	comma
	assignment
	logical
	relational
	additive
	multiplicative
	unary
	call
	member
	primary
)

type (
	statementHandler      func(p *parser) ast.Statement
	nullDenotationHandler func(p *parser) ast.Expression
	leftDenotationHandler func(p *parser, bp bindingPower, left ast.Expression) ast.Expression
)

type (
	statementLookup      map[token.TokenKind]statementHandler
	nullDenotationLookup map[token.TokenKind]nullDenotationHandler
	leftDenotationLookup map[token.TokenKind]leftDenotationHandler
	bindingPowerLookup   map[token.TokenKind]bindingPower
)

var (
	statementLookupTable      = statementLookup{}
	nullDenotationLookupTable = nullDenotationLookup{}
	leftDenotationLookupTable = leftDenotationLookup{}
	bindingPowerLookupTable   = bindingPowerLookup{}
)

func statement(kind token.TokenKind, bp bindingPower, stmt statementHandler) {
	bindingPowerLookupTable[kind] = bp
	statementLookupTable[kind] = stmt
}

func leftDenotation(kind token.TokenKind, bp bindingPower, led leftDenotationHandler) {
	bindingPowerLookupTable[kind] = bp
	leftDenotationLookupTable[kind] = led
}

func nullDenotation(kind token.TokenKind, bp bindingPower, nud nullDenotationHandler) {
	bindingPowerLookupTable[kind] = bp
	nullDenotationLookupTable[kind] = nud
}

func newTokenLookupTables() {
	leftDenotation(token.Plus, additive, parseBinaryExpression)
	leftDenotation(token.Hyphen, additive, parseBinaryExpression)

	leftDenotation(token.Star, multiplicative, parseBinaryExpression)
	leftDenotation(token.Slash, multiplicative, parseBinaryExpression)
	leftDenotation(token.Percent, multiplicative, parseBinaryExpression)
	leftDenotation(token.Hashtag, multiplicative, parseBinaryExpression)

	nullDenotation(token.Number, primary, parsePrimaryExpression)
}
