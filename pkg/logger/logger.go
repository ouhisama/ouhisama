package logger

type (
	errorCode          uint
	TokeniserErrorCode errorCode
	ParserErrorCode    errorCode
)

const (
	UnrecongnisedToken TokeniserErrorCode = iota

	UnexpectedToken ParserErrorCode = iota
	NotExpectedToken
	NoNudHandler
	NoLedHandler
)
