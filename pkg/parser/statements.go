package parser

import (
	"os"

	"github.com/ouhisama/ouhisama/pkg/ast"
	"github.com/ouhisama/ouhisama/pkg/token"
)

func parseStatement(p *parser) ast.Statement {
	if p.at().Kind == token.Newline {
		p.advance()
	}
	stmtHandler, found := statementLookupTable[p.at().Kind]
	if found {
		return stmtHandler(p)
	}

	var level uint
	if p.at().Kind == token.Indentation {
		if p.position == 0 {
			p.error(invalidIndentation, "Indentation can't be used at the top level of code", "remove the invalid indentation", "")
			os.Exit(1)
		}
		level = p.at().Length
	}

	expression := parseExpression(p, zero)
	p.want(token.Newline)
	return ast.ExpressionStatement{
		Body:  expression,
		Level: level,
	}
}
