package main

import (
	"fmt"

	"./scoping"
)

func main() {
	fmt.Println("Hello World")
	scoping.Foo()
	fmt.Println(scoping.GetScopedData())
}
