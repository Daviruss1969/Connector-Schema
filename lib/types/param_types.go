package types

type ParamType int32

const (
	PARAM_TYPE_UNKNOWN ParamType = iota
	PARAM_TYPE_URL
	PARAM_TYPE_STRING
	PARAM_TYPE_BOOL
	PARAM_TYPE_FILE
	PARAM_TYPE_SET
	PARAM_TYPE_PASSWORD
)

func (t ParamType) String() string {
	switch t {
	case PARAM_TYPE_UNKNOWN:
		return "UNKNOWN"
	case PARAM_TYPE_URL:
		return "URL"
	case PARAM_TYPE_STRING:
		return "STRING"
	case PARAM_TYPE_BOOL:
		return "BOOL"
	case PARAM_TYPE_FILE:
		return "FILE"
	case PARAM_TYPE_SET:
		return "SET"
	case PARAM_TYPE_PASSWORD:
		return "PASSWORD"
	default:
		return "UNKNOWN"
	}
}

func ToParamType(t string) ParamType {
	switch t {
	case "URL":
		return PARAM_TYPE_URL
	case "STRING":
		return PARAM_TYPE_STRING
	case "BOOL", "BOOLEAN":
		return PARAM_TYPE_BOOL
	case "FILE":
		return PARAM_TYPE_FILE
	case "SET":
		return PARAM_TYPE_SET
	case "PASSWORD":
		return PARAM_TYPE_PASSWORD
	default:
		return PARAM_TYPE_UNKNOWN
	}
}
