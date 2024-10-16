package tokeniser

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func Test_Tokeniser(t *testing.T) {
	source := "(-16 * 5 / -0.8) - (121 # 4 + 22 % 7)\n"
	tokens := Tokenise(source)
	want := []string{
		"LBracket",
		"Number: \"-16\"",
		"Whitespace",
		"Star",
		"Whitespace",
		"Number: \"5\"",
		"Whitespace",
		"Slash",
		"Whitespace",
		"Number: \"-0.8\"",
		"RBracket",
		"Whitespace",
		"Hyphen",
		"Whitespace",
		"LBracket",
		"Number: \"121\"",
		"Whitespace",
		"Hashtag",
		"Whitespace",
		"Number: \"4\"",
		"Whitespace",
		"Plus",
		"Whitespace",
		"Number: \"22\"",
		"Whitespace",
		"Percent",
		"Whitespace",
		"Number: \"7\"",
		"RBracket",
		"Newline",
		"EOF",
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
}
