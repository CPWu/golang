package main

import (
	"log"

	"github.com/cpwu/channels/helpers"
)

const numPool = 10

func CalculateValue(intChannel chan int) {
	randomNumber := helpers.RandomNumber(numPool)
	intChannel <- randomNumber
}

func main() {
	// Sends information from one part of the program to another part of the program
	intChannel := make(chan int)

	defer close(intChannel)

	go CalculateValue((intChannel))

	num := <-intChannel
	log.Println(num)
}
