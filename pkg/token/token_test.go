package token

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"
)

var tokens = []Token{
	New(EOF, ""),
	New(Newline, "\n"),
	New(Indentation, "\t"),
	New(Whitespace, " "),
	New(Identifier, "uynilo9"),
	New(Number, "213"),
	New(String, "hello world"),
	New(Equal, "="),
	New(Plus, "+"),
	New(Hyphen, "-"),
	New(Star, "*"),
	New(Slash, "/"),
	New(Percent, "%"),
	New(Hashtag, "#"),
	New(LBracket, "("),
	New(RBracket, ")"),
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
		"Hashtag",
		"LBracket",
		"RBracket",
	}
	for i, token := range tokens {
		stdout := os.Stdout
		var buffer bytes.Buffer
		read, write, _ := os.Pipe()
		os.Stdout = write
		fmt.Print(token.Debug())
		write.Close()
		os.Stdout = stdout
		io.Copy(&buffer, read)
		if buffer.String() != want[i] {
			t.Fatalf("ERROR Expected `%v`, but got `%v`", want[i], buffer.String())
		}
	}
}

func Test_New(t *testing.T) {
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
		{Kind: Hashtag, Value: "#"},
		{Kind: LBracket, Value: "("},
		{Kind: RBracket, Value: ")"},
	}
	for i, token := range tokens {
		if token != want[i] {
			t.Fatalf("ERROR Expected `%v: \"%v\"`, but got `%v \"%v\"`", want[i].Kind.string(), want[i].Value, token.Kind.string(), token.Value)
		}
	}
}
