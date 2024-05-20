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
	col "github.com/craterdog/go-collection-framework/v4"
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
	var notation = col.CDCN()
	var source = string(bytes)
	var collection = notation.ParseSource(source)

	// No validation currently required.

	// Reformat the collection file.
	source = notation.FormatCollection(collection)
	bytes = []byte(source)
	err = osx.WriteFile(collectionFile, bytes, 0644)
	if err != nil {
		panic(err)
	}
}
