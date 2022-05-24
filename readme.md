# Validate

Tired of super long long struct tags? Use this instead

## Install

```shell
go get -u github.com/vinshop/validate
```

## Usage

```go
package main

import . "github.com/vinshop/validate"

type TestStruct struct {
	A string
	B int64
}

func main() {
	data := TestStruct{
		A: "abcde",
		B: 100,
	}

	if err := Use(data, Struct(
		WithKey("some key here"),
		Register("A", Require, MinLength(5)),
		Register("B", Require),
	)).Validate(); err != nil {
		panic(err)
	}
}
```

## Document

```
constructing...
```