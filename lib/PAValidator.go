package lib

import (
	"PAValidator/lib/errors"
	"PAValidator/lib/formats"
	"PAValidator/lib/types"

	"github.com/santhosh-tekuri/jsonschema/v5"
)

func Validate(schemaPath string, data interface{}) bool {

	c := jsonschema.NewCompiler()
	c.AssertFormat = true

	// add format checkers
	c.Formats["type-format"] = formats.TypeFormat

	// create schema
	schema, err := c.Compile(schemaPath)
	if err != nil {
		errors.NewError(types.ERROR_TYPE_LEXICAL, "Schema file not in json format").Throw()
	}

	// validate input file with the schema
	if err := schema.Validate(data); err != nil {
		if verr, ok := err.(*jsonschema.ValidationError); ok {
			var msg string = ""
			for i := 0; i < len(verr.BasicOutput().Errors); i++ {
				msg = msg + verr.BasicOutput().Errors[i].KeywordLocation + " " + verr.BasicOutput().Errors[i].Error + "\n"
			}
			errors.NewError(types.ERROR_TYPE_SYNTAX, msg).Throw()
		}
	}

	return true
}
