package types

type ErrorType int32

const (
	ERROR_TYPE_UNKNOWN ErrorType = iota
	ERROR_TYPE_FLAGS
	ERROR_TYPE_PATH
	ERROR_TYPE_LEXICAL
	ERROR_TYPE_SYNTAX
	ERROR_TYPE_SCHEMA
)

func (t ErrorType) String() string {
	switch t {
	case ERROR_TYPE_UNKNOWN:
		return "ERROR"
	case ERROR_TYPE_FLAGS:
		return "FLAGS ERROR"
	case ERROR_TYPE_PATH:
		return "PATH ERROR"
	case ERROR_TYPE_LEXICAL:
		return "LEXICAL ERROR"
	case ERROR_TYPE_SYNTAX:
		return "SYNTAX ERROR"
	case ERROR_TYPE_SCHEMA:
		return "SCHEMA ERROR"
	default:
		return "UNKNOWN"
	}
}
