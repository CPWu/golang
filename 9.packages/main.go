package main

import (
	"log"

	"github.com/cpwu/mypackage/helpers"
)

// Concepts
// Packages

func main() {
	log.Println("Hello")

	var myVar helpers.SomeType
	myVar.TypeName = "Some Name"
	log.Println(myVar.TypeName)
}
