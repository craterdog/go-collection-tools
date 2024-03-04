<img src="https://craterdog.com/images/CraterDog.png" width="50%">

## Go Collection Tools

### Overview
This project provides a set of Go based tools that can validate and format a
Go collection defined using Crater Dog Collection Notation™ (aka CDCN).  For
full details on the Go Collection Framework click
[here](https://github.com/craterdog/go-collection-framework/wiki)

### Getting Started
To install the collection tools do the following from a terminal window:
```
$ git clone git@github.com:craterdog/go-collection-tools.git
$ cd go-collection-tools/
$ ./etc/build.sh
$ ls
LICENSE		bin		go.mod		src
README.md	etc		go.sum

$ls bin/
format		validate
```

### Using the Tools
The `format` command reads in a collection file formatted using Crater Dog
Collection Notation™ (aka CDCN) and reformats it in its canonical format as
follows:
```
$ go-collection-tools/bin/format example/customers.cdcn
```

The `validate` command reads in a collection file and ensures that it is
formatted using Crater Dog Collection Notation™ (aka CDCN) as follows:
```
$ go-collection-tools/bin/validate example/customers.cdcn
```

<H5 align="center"> Copyright © 2009 - 2024  Crater Dog Technologies™. All rights reserved. </H5>
