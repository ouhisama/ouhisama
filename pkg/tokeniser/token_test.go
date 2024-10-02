package tokeniser

import "testing"

var tokens = TokenList{
	New(EOF, "\x00"),
	New(Newline, "\n"),
	New(Whitespace, " "),
	New(Indentation, "\t"),
	New(Identifier, "uynilo9"),
	New(Number, "213"),
	New(String, "hello world"),
	New(Equal, "="),
	New(Plus, "+"),
	New(Hyphen, "-"),
	New(Star, "*"),
	New(Slash, "/"),
	New(Percent, "%"),
	New(LBraket, "("),
	New(RBraket, ")"),
}

func TestTokenKindString(t *testing.T) {
	want := []string{
		"EOF",
		"Newline",
		"Whitespace",
		"Indentation",
		"Identifier",
		"Number",
		"String",
		"Equal",
		"Plus",
		"Hyphen",
		"Star",
		"Slash",
		"Percent",
		"LBracket",
		"RBracket",
	}
	for i, token := range tokens {
		if token.Kind.String() != want[i] {
			t.Fatalf("ERROR Expected `%v`, got `%v`", want[i], token.Kind.String())
		}
	}
	t.Logf("SUCCESS All passed")
}

func TestTokenDebug(t *testing.T) {
	for _, token := range tokens {
		token.Debug()
	}
	// Output:
	// EOF: ""
	// Newline: ""
	// Whitespace: ""
	// Indentation: ""
	// Identifier: "uynilo9"
	// Number: "213"
	// String: "hello world"
	// Equal: ""
	// Plus: ""
	// Hyphen: ""
	// Star: ""
	// Slash: ""
	// Percent: ""
	// LBracket: ""
	// RBracket: ""
}

func TestNew(t *testing.T) {
	want := TokenList{
		{Kind: EOF, Value: "\x00"},
		{Kind: Newline, Value: "\n"},
		{Kind: Whitespace, Value: " "},
		{Kind: Indentation, Value: "\t"},
		{Kind: Identifier, Value: "uynilo9"},
		{Kind: Number, Value: "213"},
		{Kind: String, Value: "hello world"},
		{Kind: Equal, Value: "="},
		{Kind: Plus, Value: "+"},
		{Kind: Hyphen, Value: "-"},
		{Kind: Star, Value: "*"},
		{Kind: Slash, Value: "/"},
		{Kind: Percent, Value: "%"},
		{Kind: LBraket, Value: "("},
		{Kind: RBraket, Value: ")"},
	}
	for i, token := range tokens {
		if token != want[i] {
			t.Fatalf("ERROR Expected `%v: \"%v\"`, got `%v \"%v\"`", want[i].Kind.String(), want[i].Value, token.Kind.String(), token.Value)
		}
	}
	t.Logf("SUCCESS All passed")
}