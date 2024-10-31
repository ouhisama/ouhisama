package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/ouhisama/ouhisama/pkg/parser"
	"github.com/ouhisama/ouhisama/pkg/tokeniser"
	"github.com/sanity-io/litter"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		log.Fatalf("ERROR Expected 1 argument, but got %v\n", len(args))
	}

	path, err := filepath.Abs(args[1])
	if err != nil {
		log.Fatalf("ERROR Failed to get absolute path of file `%v`\n", args[1])
	}

	source, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("ERROR Failed to read file `%v`\n", path)
	}

	tokens := tokeniser.Tokenise(string(source))
	ast := parser.Parse(tokens)
	litter.Dump(ast)
}
