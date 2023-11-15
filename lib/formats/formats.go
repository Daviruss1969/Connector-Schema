package formats

import (
	"PAValidator/lib/types"
)

var TypeFormat func(interface{}) bool = func(i interface{}) bool {
	asString, ok := i.(string)
	if !ok {
		return false
	}

	return types.ToParamType(asString) != types.PARAM_TYPE_UNKNOWN
}
