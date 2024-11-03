package parser

import (
	"fmt"
	"os"
	"strconv"

	"github.com/alecthomas/colour"
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
		t := p.at()
		value, err := strconv.ParseFloat(string(t.Value), 64)
		if err != nil {
			// Program mustn't get here

			// msg := fmt.Sprintf("Failed to convert the type of `%v` unexpectedly while parsing a primary expression", t.Value)
			// advice := "We don't know why it couldn't be parsed to a float"
			// p.error(cannotParseFloat, msg, advice, err.Error())
			os.Exit(1)
		}
		p.advance()
		return ast.NumberExpression{
			Value: value,
		}
	default:
		// Program mustn't get here
		os.Exit(1)
		return nil
	}
}

func parseGroupingExpression(p *parser) ast.Expression {
	p.advance()
	expression := parseExpression(p, zero)
	p.want(token.RBracket)
	return expression
}

func parseExpression(p *parser, bp bindingPower) ast.Expression {
	t := p.at()
	nudHandler, found := nullDenotationLookupTable[t.Kind]
	if !found {
		msg := fmt.Sprintf("No null denotation handler for the token `%v`", t.Value)
		advice := colour.Sprintf("you might put an incorrect stuff ^Slike you^R^1 here")
		p.error(noNudHandler, msg, advice, "")
		os.Exit(1)
	}

	left := nudHandler(p)
	for bindingPowerLookupTable[p.at().Kind] > bp {
		t := p.at()
		ledHandler, found := leftDenotationLookupTable[t.Kind]
		if !found {
			msg := fmt.Sprintf("No left denotation handler for the token `%v`", t.Value)
			advice := colour.Sprintf("you might put an incorrect stuff ^Slike you^R^1 here")
			p.error(noLedHandler, msg, advice, "")
			os.Exit(1)
		}
		left = ledHandler(p, bindingPowerLookupTable[p.at().Kind], left)
	}

	return left
}
