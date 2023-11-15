package types

type ErrorType int32

const (
	ERROR_TYPE_UNKNOWN ErrorType = iota
	ERROR_TYPE_PATH
	ERROR_TYPE_LEXICAL
	ERROR_TYPE_SYNTAX
)

func (t ErrorType) String() string {
	switch t {
	case ERROR_TYPE_UNKNOWN:
		return "ERROR"
	case ERROR_TYPE_PATH:
		return "PATH ERROR"
	case ERROR_TYPE_LEXICAL:
		return "LEXICAL ERROR"
	case ERROR_TYPE_SYNTAX:
		return "SYNTAX ERROR"
	default:
		return "UNKNOWN"
	}
}
