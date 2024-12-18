package helpers

import (
	"math/rand"
	"time"
)

func RandomNumber(n int) int {
	// Seeding with time otherwise the return of Rand is the same everytime.
	rand.Seed(time.Now().UnixNano())
	// Using the built-in Golang package `math`
	// Leverage Rand to generate an INT given a pool size
	value := rand.Intn(n)

	return value
}
