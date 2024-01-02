package lib

import (
	"ConnectorSchema/lib/errors"
	"ConnectorSchema/lib/types"

	"github.com/santhosh-tekuri/jsonschema/v5"
)

func Validate(schemaPath string, data interface{}) *errors.Error {
	c := jsonschema.NewCompiler()

	// create schema
	schema, err := c.Compile(schemaPath)
	if err != nil {
		if serr, ok := err.(*jsonschema.SchemaError); ok {
			return errors.NewError(types.ERROR_TYPE_SCHEMA, serr.Error())
		}
	}

	// validate input file with the schema
	if err := schema.Validate(data); err != nil {
		if verr, ok := err.(*jsonschema.ValidationError); ok {
			var msg string = "\n"
			for i := 0; i < len(verr.BasicOutput().Errors); i++ {
				msg = msg + verr.BasicOutput().Errors[i].KeywordLocation + " " + verr.BasicOutput().Errors[i].Error + "\n"
			}
			return errors.NewError(types.ERROR_TYPE_SYNTAX, msg)
		}
	}

	return nil
}
