package lib

import (
	"ConnectorSchema/lib/errors"
	"ConnectorSchema/lib/types"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"testing"
)

var schemaTestPath string = "validator/schema.json"

func readInputFile(path string) (interface{}, *errors.Error) {
	inputF, err := os.Open(path)
	if err != nil {
		return nil, errors.NewError(types.ERROR_TYPE_PATH, "Input file \""+path+"\" not found")
	}

	inputS, _ := io.ReadAll(inputF)
	defer inputF.Close()

	// parse input file
	var v interface{}
	if err := json.Unmarshal([]byte(inputS), &v); err != nil {
		return nil, errors.NewError(types.ERROR_TYPE_LEXICAL, "Input file not in json format")
	}

	return v, nil
}

func TestInvalid00emptyFile(t *testing.T) {
	var inputPath string = "test/invalid/00-emptyFile.json"

	_, err := readInputFile(inputPath)
	if err == nil {
		t.Fail()
	}
}

func TestInvalid01badjson(t *testing.T) {
	var inputPath string = "test/invalid/00-badjson"

	_, err := readInputFile(inputPath)
	if err == nil {
		t.Fail()
	}
}

func TestInvalid02emptyJSON(t *testing.T) {
	var inputPath string = "test/invalid/02-emptyJson.json"

	data, err := readInputFile(inputPath)
	if err != nil {
		fmt.Println("ici")
		t.Fail()
	}

	libErr := Validate(data, schemaTestPath)
	if libErr == nil {
		t.Fail()
	} else {
		if libErr.Type != types.ERROR_TYPE_SYNTAX {
			t.Fail()
		}
	}
}

func TestInvalid03missingTopProp(t *testing.T) {
	var inputPath string = "test/invalid/03-missingTopProp.json"

	data, err := readInputFile(inputPath)
	if err != nil {
		t.Fail()
	}

	libErr := Validate(data, schemaTestPath)
	if libErr == nil {
		t.Fail()
	} else {
		if libErr.Type != types.ERROR_TYPE_SYNTAX {
			t.Fail()
		}
	}
}

func TestInvalid04wrongTypeTopProp(t *testing.T) {
	var inputPath string = "test/invalid/04-wrongTypeTopProp.json"

	data, err := readInputFile(inputPath)
	if err != nil {
		t.Fail()
	}

	libErr := Validate(data, schemaTestPath)
	if libErr == nil {
		t.Fail()
	} else {
		if libErr.Type != types.ERROR_TYPE_SYNTAX {
			t.Fail()
		}
	}
}

func TestInvalid05missingNestedProp(t *testing.T) {
	var inputPath string = "test/invalid/05-missingNestedProp.json"

	data, err := readInputFile(inputPath)
	if err != nil {
		t.Fail()
	}

	libErr := Validate(data, schemaTestPath)
	if libErr == nil {
		t.Fail()
	} else {
		if libErr.Type != types.ERROR_TYPE_SYNTAX {
			t.Fail()
		}
	}
}

func TestInvalid06wrongTypeNestedProp(t *testing.T) {
	var inputPath string = "test/invalid/06-wrongTypeNestedProp.json"

	data, err := readInputFile(inputPath)
	if err != nil {
		t.Fail()
	}

	libErr := Validate(data, schemaTestPath)
	if libErr == nil {
		t.Fail()
	} else {
		if libErr.Type != types.ERROR_TYPE_SYNTAX {
			t.Fail()
		}
	}
}

func TestInvalid07invalidTypeFormat(t *testing.T) {
	var inputPath string = "test/invalid/07-invalidTypeFormat.json"

	data, err := readInputFile(inputPath)
	if err != nil {
		t.Fail()
	}

	libErr := Validate(data, schemaTestPath)
	if libErr == nil {
		t.Fail()
	} else {
		if libErr.Type != types.ERROR_TYPE_SYNTAX {
			t.Fail()
		}
	}
}

func TestInvalid08invalidMinimumOID(t *testing.T) {
	var inputPath string = "test/invalid/08-invalidMinimumOID.json"

	data, err := readInputFile(inputPath)
	if err != nil {
		t.Fail()
	}

	libErr := Validate(data, schemaTestPath)
	if libErr == nil {
		t.Fail()
	} else {
		if libErr.Type != types.ERROR_TYPE_SYNTAX {
			t.Fail()
		}
	}
}

func TestInvalid09invalidschemaTestPath(t *testing.T) {
	var inputPath string = "test/invalid/09-invalidSchemaPath.json"

	data, err := readInputFile(inputPath)
	if err != nil {
		t.Fail()
	}

	libErr := Validate(data, "invalid/schemaTestPath")
	if libErr == nil {
		t.Fail()
	} else {
		if libErr.Type != types.ERROR_TYPE_SCHEMA {
			t.Fail()
		}
	}
}

func TestInvalid10invalidSchemaFile(t *testing.T) {
	libErr := Validate(nil, "test/invalid/10-invalidSchemaFile.jsd")
	if libErr == nil {
		t.Fail()
	} else {
		if libErr.Type != types.ERROR_TYPE_SCHEMA {
			t.Fail()
		}
	}
}

func TestInvalid11invalidCSTINFO(t *testing.T) {
	var inputPath string = "test/invalid/11-invalidCSTINFO.json"

	data, err := readInputFile(inputPath)
	if err != nil {
		t.Fail()
	}

	libErr := Validate(data, schemaTestPath)
	if libErr == nil {
		t.Fail()
	} else {
		if libErr.Type != types.ERROR_TYPE_SYNTAX {
			t.Fail()
		}
	}
}

func TestInvalid12invalidCSTINFOformat(t *testing.T) {
	var inputPath string = "test/invalid/12-invalidCSTINFOformat.json"

	data, err := readInputFile(inputPath)
	if err != nil {
		t.Fail()
	}

	libErr := Validate(data, schemaTestPath)
	if libErr == nil {
		t.Fail()
	} else {
		if libErr.Type != types.ERROR_TYPE_SYNTAX {
			t.Fail()
		}
	}
}

func TestInvalid13invalidCSTINFOenum(t *testing.T) {
	var inputPath string = "test/invalid/13-invalidCSTINFOenum.json"

	data, err := readInputFile(inputPath)
	if err != nil {
		t.Fail()
	}

	libErr := Validate(data, schemaTestPath)
	if libErr == nil {
		t.Fail()
	} else {
		if libErr.Type != types.ERROR_TYPE_SYNTAX {
			t.Fail()
		}
	}
}

func TestValid00validRFX(t *testing.T) {
	var inputPath string = "test/valid/00-validRFX.json"

	data, err := readInputFile(inputPath)
	if err != nil {
		t.Fail()
	}

	libErr := Validate(data, schemaTestPath)
	if libErr != nil {
		t.Fail()
	}
}

func TestValid01validTDI(t *testing.T) {
	var inputPath string = "test/valid/01-validTDI.json"

	data, err := readInputFile(inputPath)
	if err != nil {
		t.Fail()
	}

	libErr := Validate(data, schemaTestPath)
	if libErr != nil {
		t.Fail()
	}
}
