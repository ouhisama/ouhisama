package token

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
	Hashtag

	LBracket
	RBracket
)

func (k TokenKind) String() string {
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
	case Hashtag:
		return "Hashtag"
	case LBracket:
		return "LBracket"
	case RBracket:
		return "RBracket"
	default:
		return "Unknown"
	}
}

type TokenPosition struct {
	Index  uint
	Column uint
	Line   uint
}

type Token struct {
	Kind     TokenKind
	Value    string
	Position TokenPosition
	Length   uint
}

func (t Token) isOneOf(kinds ...TokenKind) bool {
	return slices.Contains(kinds, t.Kind)
}

func (t Token) Debug() string {
	if t.isOneOf(Identifier, Number, String) {
		return fmt.Sprintf("\nKind: %v,\nValue: \"%v\",\nPosition:\n\tIndex: %v,\n\tColumn: %v,\n\tLine: %v\nLength: %v\n", t.Kind.String(), t.Value, t.Position.Index, t.Position.Column, t.Position.Line, t.Length)
	} else {
		return fmt.Sprintf("\nKind: %v,\nValue: \"%v\",\nPosition:\n\tIndex: %v,\n\tColumn: %v,\n\tLine: %v\nLength: %v\n", t.Kind.String(), "", t.Position.Index, t.Position.Column, t.Position.Line, t.Length)
	}
}

func NewToken(kind TokenKind, value string, position TokenPosition, length uint) Token {
	return Token{
		Kind:     kind,
		Value:    value,
		Position: position,
		Length:   length,
	}
}

func NewTokenPosition(index uint, column uint, line uint) TokenPosition {
	return TokenPosition{
		Index:  index,
		Column: column,
		Line:   line,
	}
}
