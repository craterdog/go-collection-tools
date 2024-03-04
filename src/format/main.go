package main

import (
	fmt "fmt"
	col "github.com/craterdog/go-collection-framework/v3"
	cdc "github.com/craterdog/go-collection-framework/v3/cdcn"
	osx "os"
)

// MAIN PROGRAM

func main() {
	// Validate the commandline arguments.
	if len(osx.Args) < 2 {
		fmt.Println("Usage: format <collection-file>")
		return
	}
	var collectionFile = osx.Args[1]

	// Parse the collection file.
	var bytes, err = osx.ReadFile(collectionFile)
	if err != nil {
		panic(err)
	}
	var source = string(bytes)
	var parser = cdc.Parser().Make()
	var collection = parser.ParseSource(source)

	// Reformat the collection file.
	var formatter = col.Formatter().Make()
	source = formatter.FormatCollection(collection)
	bytes = []byte(source)
	err = osx.WriteFile(collectionFile, bytes, 0644)
	if err != nil {
		panic(err)
	}
}
