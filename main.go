package main

import (
	"PAValidator/lib"
	"PAValidator/lib/errors"
	"PAValidator/lib/types"
	"encoding/json"
	"flag"
	"io"
	"os"
)

func main() {
	inputFile := flag.String("i", "", "Input file")
	flag.Parse()

	inputF, err := os.Open(*inputFile)
	if err != nil {
		errors.NewError(types.ERROR_TYPE_PATH, "Input file : "+*inputFile+" not found").Throw()
	}
	inputS, _ := io.ReadAll(inputF)
	defer inputF.Close()

	var v interface{}
	if err := json.Unmarshal([]byte(inputS), &v); err != nil {
		errors.NewError(types.ERROR_TYPE_LEXICAL, "Input file not in json format").Throw()
	}

	lib.Validate("lib/validator/schema.json", v)
}
