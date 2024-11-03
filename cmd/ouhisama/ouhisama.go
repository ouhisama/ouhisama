package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/ouhisama/ouhisama/pkg/logger"
	"github.com/ouhisama/ouhisama/pkg/parser"
	"github.com/ouhisama/ouhisama/pkg/tokeniser"
	"github.com/sanity-io/litter"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		msg := fmt.Sprintf("Expected 1 argument, but got %v", len(args))
		logger.Error(logger.WrongArgumentCount, msg, "")
		os.Exit(1)
	}

	path, err := filepath.Abs(args[1])
	if err != nil {
		msg := fmt.Sprintf("Failed to get absolute path of file `%v`", args[1])
		logger.Error(logger.CannotGetAbsolutePath, msg, err.Error())
		os.Exit(1)
	}

	source, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			msg := fmt.Sprintf("The path `%v` you provided doesn't exist", path)
			logger.Error(logger.InputPathFileNotFound, msg, "")
			os.Exit(1)
		}
		msg := fmt.Sprintf("Failed to read file `%v` for some reason", path)
		logger.Error(logger.CannotReadInputPathFile, msg, err.Error())
		os.Exit(1)
	}

	tokens := tokeniser.Tokenise(path, string(source))
	// for _, token := range tokens {
	// 	fmt.Println(token.Debug())
	// }
	ast := parser.Parse(path, string(source), tokens)
	litter.Dump(ast)
}
