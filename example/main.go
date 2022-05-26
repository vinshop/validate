package main

import (
	. "github.com/vinshop/validate"
	"log"
)

type TestStruct struct {
	A interface{}
	B int64
	C []TestStruct
}

func main() {
	data := TestStruct{
		A: "abcde",
		B: 100,
		C: []TestStruct{
			{
				A: "Hello",
				B: 123,
			},
			{
				A: 123,
				B: 123,
			},
		},
	}

	if err := Use(data, Struct(
		WithKey("some key here"),
		Field("A", Require, MinLength(5)),
		Field("B", Require),
		Field("C", Array(
			Each(
				Struct(
					Field("A", String()),
				),
			),
		)),
	)).Validate(); err != nil {
		log.Fatal(err)
	}
}
