/*
................................................................................
.    Copyright (c) 2009-2024 Crater Dog Technologies.  All Rights Reserved.    .
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See https://opensource.org/license/MIT)                        .
................................................................................
*/

package main

import (
	fmt "fmt"
	cdc "github.com/craterdog/go-collection-framework/v4/cdcn"
	col "github.com/craterdog/go-collection-framework/v4/collection"
	osx "os"
)

func main() {
	var collectionFile = retrieveArguments()
	var collection = parseCollection(collectionFile)
	validateCollection(collection)
	reformatCollection(collectionFile, collection)
}

func retrieveArguments() (collectionFile string) {
	if len(osx.Args) < 2 {
		fmt.Println("Usage: format <collection-file>")
		osx.Exit(1)
	}
	collectionFile = osx.Args[1]
	return collectionFile
}

func parseCollection(collectionFile string) col.Collection {
	var bytes, err = osx.ReadFile(collectionFile)
	if err != nil {
		panic(err)
	}
	var source = string(bytes)
	var parser = cdc.Parser().Make()
	var collection = parser.ParseSource(source)
	return collection
}

func validateCollection(collection col.Collection) {
	// No validation currently required.
}

func reformatCollection(
	collectionFile string,
	collection col.Collection,
) {
	var formatter = cdc.Formatter().Make()
	var source = formatter.FormatCollection(collection)
	var bytes = []byte(source)
	var err = osx.WriteFile(collectionFile, bytes, 0644)
	if err != nil {
		panic(err)
	}
}
