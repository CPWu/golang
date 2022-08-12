package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")

	var whatToSay string
	var integer int

	whatToSay = "Goodbye, cruel world."
	fmt.Println(whatToSay)

	integer = 10
	fmt.Println("integer is set to:",integer)

	whatWasSaid := saySomething()
	fmt.Println("The funcion returned:", whatWasSaid)
}

func saySomething() string {
	return "something"
}