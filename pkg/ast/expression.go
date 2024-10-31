package ast

import "github.com/ouhisama/ouhisama/pkg/token"

type NumberExpression struct {
	Value float64
}

func (_ NumberExpression) expression() {}

type BinaryExpression struct {
	Operator token.Token
	Left     Expression
	Right    Expression
}

func (_ BinaryExpression) expression() {}
