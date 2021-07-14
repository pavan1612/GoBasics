package main

import (
	"fmt"
)

type A struct {
	field1 int
	field2 string
}

// Creating a struct B with fields of A
type B struct {
	A
	field3 bool
	field4 float32
}

func main() {
	// B works like is inheriting the fields/properties from A
	b := B{
		A: A{
			field1: 1,
			field2: "1",
		},
		field3: true,
		field4: 1.0,
	}
	fmt.Println(b)

	// Second and easier way of creating an object
	bb := B{}
	bb.field1 = 2
	bb.field2 = "2"
	bb.field3 = false
	bb.field4 = 2.0
	fmt.Println(bb)
}
