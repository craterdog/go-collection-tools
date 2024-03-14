/*
................................................................................
.    Copyright (c) 2009-2024 Crater Dog Technologies™.  All Rights Reserved.   .
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See http://opensource.org/licenses/MIT)                        .
................................................................................
*/

package main

import (
	fmt "fmt"
	cdc "github.com/craterdog/go-collection-framework/v3/cdcn"
	osx "os"
)

// MAIN PROGRAM

func main() {
	// Validate the commandline arguments.
	if len(osx.Args) < 2 {
		fmt.Println("Usage: validate <collection-file>")
		return
	}
	var collectionFile = osx.Args[1]

	// Validate the collection file.
	var bytes, err = osx.ReadFile(collectionFile)
	if err != nil {
		panic(err)
	}
	var source = string(bytes)
	var parser = cdc.Parser().Make()
	parser.ParseSource(source)
}
