package parser

import (
	"fmt"
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
			log.Fatalf("ERROR Failed to convert type of the number `%v` unexpectedly while parsing\n", t.Value)
		}
		fmt.Println(value)
		return ast.NumberExpression{
			Value: value,
		}
	default:
		log.Fatalf("ERROR Unexpected token `%v` while creating a primary expression in the parser\n", p.at().Kind.String())
		return nil
	}
}

func parseExpression(p *parser, bp bindingPower) ast.Expression {
	t := p.at()
	nudHandler, found := nullDenotationLookupTable[t.Kind]
	if !found {
		log.Fatalf("ERROR No null denotation handler for the token `%v` (having the type `%v`) in the parser\n", t.Value, t.Kind.String())
	}

	left := nudHandler(p)
	for bindingPowerLookupTable[p.at().Kind] > bp {
		t := p.at()
		ledHandler, found := leftDenotationLookupTable[t.Kind]
		if !found {
			log.Fatalf("ERROR No left denotation handler for the token `%v` (having the type `%v`) in the parser\n", t.Value, t.Kind.String())
		}
		left = ledHandler(p, bindingPowerLookupTable[p.at().Kind], left)
	}

	return left
}
