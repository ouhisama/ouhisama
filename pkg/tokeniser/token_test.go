package tokeniser

import "testing"

const length = 15

var tokens = [length]Token{
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

func TestTokenKind_String(t *testing.T) {
	want := [length]string{
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
			t.Fatalf("ERROR Expected `%v`, but got `%v`", want[i], token.Kind.String())
		}
	}
	t.Log("SUCCESS All passed")
}

func TestToken_isOneOf(t *testing.T) {
	want := [length]bool{
		false,
		false,
		false,
		false,
		true,
		true,
		true,
		false,
		false,
		false,
		false,
		false,
	}
	for i, token := range tokens {
		if token.isOneOf(Identifier, Number, String) != want[i] {
			t.Fatalf("ERROR Expected `%v: \"%v\"` to get `%v`, but got `%v`", token.Kind.String(), token.Value, want[i], token.isOneOf(Identifier, Number, String))
		}
	}
	t.Log("SUCCESS All passed")
}

func TestToken_Debug(t *testing.T) {
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

func Test_New(t *testing.T) {
	want := [length]Token{
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
			t.Fatalf("ERROR Expected `%v: \"%v\"`, but got `%v \"%v\"`", want[i].Kind.String(), want[i].Value, token.Kind.String(), token.Value)
		}
	}
	t.Log("SUCCESS All passed")
}
