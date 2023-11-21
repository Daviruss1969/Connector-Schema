package main

import (
	"ConnectorSchema/lib"
	"ConnectorSchema/lib/errors"
	"ConnectorSchema/lib/types"
	"encoding/json"
	"flag"
	"io"
	"os"
)

func main() {
	// get input file path
	inputFile := flag.String("i", "", "Input file")
	flag.Parse()

	// handle flags error
	if *inputFile == "" {
		errors.NewError(types.ERROR_TYPE_FLAGS, "Flag -i is needed to test an input file").Throw()
	}

	// read input file
	inputF, err := os.Open(*inputFile)
	if err != nil {
		errors.NewError(types.ERROR_TYPE_PATH, "Input file \""+*inputFile+"\" not found").Throw()
	}
	inputS, _ := io.ReadAll(inputF)
	defer inputF.Close()

	// parse input file
	var v interface{}
	if err := json.Unmarshal([]byte(inputS), &v); err != nil {
		errors.NewError(types.ERROR_TYPE_LEXICAL, "Input file not in json format").Throw()
	}

	libErr := lib.Validate("lib/validator/schema.json", v)
	if libErr != nil {
		libErr.Throw()
	}
}
