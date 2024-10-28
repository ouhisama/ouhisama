package ast

type BlockStatement struct {
	Body []Statement
}

func (_ BlockStatement) statement() {}

type ExpressionStatement struct {
	Body Expression
}

func (_ ExpressionStatement) statement() {}
