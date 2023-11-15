package main

import (
	"PAValidator/lib"
	"flag"
	"io"
	"os"
)

func main() {
	inputFile := flag.String("i", "", "Input file")
	flag.Parse()

	inputF, _ := os.Open(*inputFile)
	inputS, _ := io.ReadAll(inputF)
	defer inputF.Close()

	schemaF, _ := os.Open("lib/validator/schema.json")
	schemaS, _ := io.ReadAll(schemaF)
	defer schemaF.Close()

	lib.Validate(string(inputS), string(schemaS))
}
