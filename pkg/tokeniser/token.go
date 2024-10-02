package tokeniser

import "fmt"

type TokenKind uint

const (
	EOF TokenKind = iota
	Newline
	Whitespace
	Indentation

	Identifier
	Number
	String

	Equal

	Plus
	Hyphen
	Star
	Slash
	Percent

	LBraket
	RBraket
)

func (kind TokenKind) String() string {
	switch kind {
	case EOF:
		return "EOF"
	case Newline:
		return "Newline"
	case Whitespace:
		return "Whitespace"
	case Indentation:
		return "Indentation"
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
	case LBraket:
		return "LBracket"
	case RBraket:
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

func (token Token) isOneOf(kinds ...TokenKind) bool {
	for _, kind := range kinds {
		if token.Kind == kind {
			return true
		}
	}
	return false
}

func (token Token) Debug() {
	if token.isOneOf(Identifier, Number, String) {
		fmt.Printf("%v: \"%v\"\n", token.Kind.String(), token.Value)
	} else {
		fmt.Printf("%v: \"\"\n", token.Kind.String())
	}
}

type TokenList []Token

func New(kind TokenKind, value TokenValue) Token {
	return Token{
		Kind:  kind,
		Value: value,
	}
}
