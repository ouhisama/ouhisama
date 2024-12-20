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

func newStatement(kind token.TokenKind, stmt statementHandler) {
	bindingPowerLookupTable[kind] = zero
	statementLookupTable[kind] = stmt
}

func newLeftDenotation(kind token.TokenKind, bp bindingPower, led leftDenotationHandler) {
	bindingPowerLookupTable[kind] = bp
	leftDenotationLookupTable[kind] = led
}

func newNullDenotation(kind token.TokenKind, nud nullDenotationHandler) {
	bindingPowerLookupTable[kind] = primary
	nullDenotationLookupTable[kind] = nud
}

func newTokenLookupTables() {
	newStatement(token.Newline, parseStatement)

	newLeftDenotation(token.Plus, additive, parseBinaryExpression)
	newLeftDenotation(token.Hyphen, additive, parseBinaryExpression)

	newLeftDenotation(token.Star, multiplicative, parseBinaryExpression)
	newLeftDenotation(token.Slash, multiplicative, parseBinaryExpression)
	newLeftDenotation(token.Percent, multiplicative, parseBinaryExpression)
	newLeftDenotation(token.Hashtag, multiplicative, parseBinaryExpression)

	newNullDenotation(token.Number, parsePrimaryExpression)
	newNullDenotation(token.LBracket, parseGroupingExpression)
}
