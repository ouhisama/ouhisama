package parser

import (
	"log"
	"strconv"

	"github.com/ouhisama/ouhisama/pkg/ast"
	"github.com/ouhisama/ouhisama/pkg/token"
)

func parseBinaryExpression(p *parser, bp bindingPower, left ast.Expression) ast.Expression {
	operator := p.eat()
	right := parseExpression(p, bp)
	return ast.BinaryExpression{
		Operator: operator,
		Left:     left,
		Right:    right,
	}
}

func parsePrimaryExpression(p *parser) ast.Expression {
	switch p.at().Kind {
	case token.Number:
		t := p.eat()
		value, err := strconv.ParseFloat(string(t.Value), 64)
		if err != nil {
			log.Fatalf("ERROR Failed to convert the value type of `%v` unexpectedly while creating a primary expression\n", t.Value)
		}
		return ast.NumberExpression{
			Value: value,
		}
	default:
		log.Fatalf("ERROR Unexpected token `%v` while creating a primary expression\n", p.at().Kind.String())
		return nil
	}
}

func parseExpression(p *parser, bp bindingPower) ast.Expression {
	t := p.at()
	nudHandler, found := nullDenotationLookupTable[t.Kind]
	if !found {
		log.Fatalf("ERROR No null denotation handler for the token `%v` while parsing an expression\n", t.Value)
	}

	left := nudHandler(p)
	for bindingPowerLookupTable[p.at().Kind] > bp {
		t := p.at()
		ledHandler, found := leftDenotationLookupTable[t.Kind]
		if !found {
			log.Fatalf("ERROR No left denotation handler for the token `%v` while parsing an expression\n", t.Value)
		}
		left = ledHandler(p, bindingPowerLookupTable[p.at().Kind], left)
	}

	return left
}
