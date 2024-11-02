package parser

import (
	"fmt"
	"os"
	"strconv"

	"github.com/ouhisama/ouhisama/pkg/ast"
	"github.com/ouhisama/ouhisama/pkg/logger"
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
		value, _ := strconv.ParseFloat(string(t.Value), 64)
		// if err != nil {
		// 	log.Fatalf("ERROR Failed to convert the value type of `%v` unexpectedly while creating a primary expression\n", t.Value)
		// }
		return ast.NumberExpression{
			Value: value,
		}
	default:
		err, code := logger.UnexpectedTokenError(p.file, p.at())
		fmt.Println(err)
		os.Exit(int(code))
		return nil
	}
}

func parseExpression(p *parser, bp bindingPower) ast.Expression {
	t := p.at()
	nudHandler, found := nullDenotationLookupTable[t.Kind]
	if !found {
		err, code := logger.NoNudHandlerError(p.file, t)
		fmt.Println(err)
		os.Exit(int(code))
	}

	left := nudHandler(p)
	for bindingPowerLookupTable[p.at().Kind] > bp {
		t := p.at()
		ledHandler, found := leftDenotationLookupTable[t.Kind]
		if !found {
			err, code := logger.NoLedHandlerError(p.file, t)
			fmt.Println(err)
			os.Exit(int(code))
		}
		left = ledHandler(p, bindingPowerLookupTable[p.at().Kind], left)
	}

	return left
}
