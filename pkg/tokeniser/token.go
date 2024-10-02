package tokeniser

import (
	"fmt"
	"slices"
)

type TokenKind uint

const (
	EOF TokenKind = iota
	Newline
	Indentation
	Whitespace

	Identifier
	Number
	String

	Equal

	Plus
	Hyphen
	Star
	Slash
	Percent

	LBracket
	RBracket
)

func (k TokenKind) string() string {
	switch k {
	case EOF:
		return "EOF"
	case Newline:
		return "Newline"
	case Indentation:
		return "Indentation"
	case Whitespace:
		return "Whitespace"
	case Identifier:
		return "Identifier"
	case Number:
		return "Number"
	case String:
		return "String"
	case Equal:
		return "Equal"
	case Plus:
		return "Plus"
	case Hyphen:
		return "Hyphen"
	case Star:
		return "Star"
	case Slash:
		return "Slash"
	case Percent:
		return "Percent"
	case LBracket:
		return "LBracket"
	case RBracket:
		return "RBracket"
	default:
		return "Unknown"
	}
}

type TokenValue string

type Token struct {
	Kind  TokenKind
	Value TokenValue
}

func (t Token) isOneOf(kinds ...TokenKind) bool {
	return slices.Contains(kinds, t.Kind)
}

func (t Token) Debug() {
	if t.isOneOf(Identifier, Number, String) {
		fmt.Printf("%v: \"%v\"\n", t.Kind.string(), t.Value)
	} else {
		fmt.Printf("%v\n", t.Kind.string())
	}
}

func newToken(kind TokenKind, value TokenValue) Token {
	return Token{
		Kind:  kind,
		Value: value,
	}
}
