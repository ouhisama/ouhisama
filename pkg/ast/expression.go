package ast

import "github.com/ouhisama/ouhisama/pkg/token"

type NumberExpression struct {
	Value float64
}

func (_ NumberExpression) expression() {}

type BinaryExpression struct {
	Left     Expression
	Operator token.Token
	Right    Expression
}

func (_ BinaryExpression) expression() {}
