package tokeniser

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func Test_Tokeniser(t *testing.T) {
	source := `(10 + 20) * 30`
	tokens := Tokenise(source)
	want := []string{
		"LBracket\n",
		"Number: \"10\"\n",
		"Whitespace\n",
		"Plus\n",
		"Whitespace\n",
		"Number: \"20\"\n",
		"RBracket\n",
		"Whitespace\n",
		"Star\n",
		"Whitespace\n",
		"Number: \"30\"\n",
		"EOF\n",
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
