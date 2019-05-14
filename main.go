package main

import (
	"fmt"
)

type Foo struct {
	Data string
}

func newFoo() *Foo {
	foo := new(Foo)
	foo.Data = "bar"
	return foo
}

func literalFoo() Foo {
	return Foo{Data: "bar"}
}

func main() {
	fmt.Printf("New foo: %v\n", newFoo())
	fmt.Printf("Literal foo: %v\n", literalFoo())
}
