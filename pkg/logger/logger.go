package logger

import (
	"fmt"

	"github.com/alecthomas/colour"
)

type errorKind uint

const (
	WrongArgumentCount errorKind = iota + 1
	CannotGetAbsolutePath
	InputPathFileNotFound
	CannotReadInputPathFile
)

func Error(kind errorKind, msg string, detail string) {
	err := fmt.Sprintf("G%0*d", 3, kind)

	if detail != "" {
		detail = "\n" + detail + "\n"
	}

	fmt.Println(colour.Sprintf("^1ERROR[%v]^7 %v\n%v", err, msg, detail))
}
