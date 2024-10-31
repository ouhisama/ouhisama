package parser

import "github.com/ouhisama/ouhisama/pkg/ast"

func parseStatement(p *parser) ast.Statement {
	stmtHandler, found := statementLookupTable[p.at().Kind]
	if found {
		return stmtHandler(p)
	}

	expression := parseExpression(p, zero)
	return ast.ExpressionStatement{
		Body: expression,
	}
}
