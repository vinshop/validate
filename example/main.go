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
		Field("A", Require, MinLength(5)),
		Field("B", Require),
	)).Validate(); err != nil {
		panic(err)
	}
}
