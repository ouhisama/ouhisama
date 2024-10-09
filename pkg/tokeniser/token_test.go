package tokeniser

import (
	"bytes"
	"io"
	"os"
	"testing"
)

var tokens = []Token{
	newToken(EOF, ""),
	newToken(Newline, "\n"),
	newToken(Indentation, "\t"),
	newToken(Whitespace, " "),
	newToken(Identifier, "uynilo9"),
	newToken(Number, "213"),
	newToken(String, "hello world"),
	newToken(Equal, "="),
	newToken(Plus, "+"),
	newToken(Hyphen, "-"),
	newToken(Star, "*"),
	newToken(Slash, "/"),
	newToken(Percent, "%"),
	newToken(LBracket, "("),
	newToken(RBracket, ")"),
}

func TestToken_Debug(t *testing.T) {
	want := []string{
		"EOF",
		"Newline",
		"Indentation",
		"Whitespace",
		"Identifier: \"uynilo9\"",
		"Number: \"213\"",
		"String: \"hello world\"",
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
		stdout := os.Stdout
		var buffer bytes.Buffer
		read, write, _ := os.Pipe()
		os.Stdout = write
		token.Debug()
		write.Close()
		os.Stdout = stdout
		io.Copy(&buffer, read)
		if buffer.String() != want[i] {
			t.Fatalf("ERROR Expected `%v`, but got `%v`", want[i], buffer.String())
		}
	}
	t.Log("SUCCESS All passed")
}

func Test_newToken(t *testing.T) {
	want := []Token{
		{Kind: EOF, Value: ""},
		{Kind: Newline, Value: "\n"},
		{Kind: Indentation, Value: "\t"},
		{Kind: Whitespace, Value: " "},
		{Kind: Identifier, Value: "uynilo9"},
		{Kind: Number, Value: "213"},
		{Kind: String, Value: "hello world"},
		{Kind: Equal, Value: "="},
		{Kind: Plus, Value: "+"},
		{Kind: Hyphen, Value: "-"},
		{Kind: Star, Value: "*"},
		{Kind: Slash, Value: "/"},
		{Kind: Percent, Value: "%"},
		{Kind: LBracket, Value: "("},
		{Kind: RBracket, Value: ")"},
	}
	for i, token := range tokens {
		if token != want[i] {
			t.Fatalf("ERROR Expected `%v: \"%v\"`, but got `%v \"%v\"`", want[i].Kind.string(), want[i].Value, token.Kind.string(), token.Value)
		}
	}
	t.Log("SUCCESS All passed")
}
