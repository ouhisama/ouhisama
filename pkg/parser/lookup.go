package parser

import (
	"strconv"

	"github.com/ouhisama/ouhisama/pkg/ast"
	"github.com/ouhisama/ouhisama/pkg/token"
)

type BindingPower uint

const (
	_Default BindingPower = iota
	_Comma
	_Assignment
	_Logical
	_Relational
	_Additive
	_Multiplicative
	_Unary
	_Call
	_Member
	_Primary
)

type (
	statementHandler      func(p *parser) ast.Statement
	nullDenotationHandler func(p *parser) ast.Expression
	leftDenotationHandler func(p *parser, left ast.Expression, bp BindingPower) ast.Expression
)

type (
	statementLookup      map[token.TokenKind]statementHandler
	nullDenotationLookup map[token.TokenKind]nullDenotationHandler
	leftDenotationLookup map[token.TokenKind]leftDenotationHandler
	bindingPowerLookup   map[token.TokenKind]BindingPower
)

var (
	statementLookupTable      = statementLookup{}
	nullDenotationLookupTable = nullDenotationLookup{}
	leftDenotationLookupTable = leftDenotationLookup{}
	bindingPowerLookupTable   = bindingPowerLookup{}
)

func statement(kind token.TokenKind, bp BindingPower, stmt statementHandler) {
	bindingPowerLookupTable[kind] = bp
	statementLookupTable[kind] = stmt
}

func leftDenotation(kind token.TokenKind, bp BindingPower, led leftDenotationHandler) {
	bindingPowerLookupTable[kind] = bp
	leftDenotationLookupTable[kind] = led
}

func nullDenotation(kind token.TokenKind, bp BindingPower, nud nullDenotationHandler) {
	bindingPowerLookupTable[kind] = bp
	nullDenotationLookupTable[kind] = nud
}

func newTokenLookups() {
	nullDenotation(token.Number, _Primary, func(p *parser) ast.Expression {
		value, err := strconv.ParseFloat(string(p.eat().Value), 64)
		if err != nil {
		}
		return ast.NumberExpression{
			Value: value,
		}
	})
}
