package main

import (
	"fmt"

	"example.com/scoping"
)

func main() {
	fmt.Println("Hello World")
	scoping.Foo()
	fmt.Println(scoping.GetScopedData())
}
