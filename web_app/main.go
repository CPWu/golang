package main

import (
	"errors"
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// Home is the home page handler.
func Home(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Home Page")
}

// About is the about page handler.
func About(w http.ResponseWriter, r *http.Request) {
	sum := AddValues(2, 2)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("This is at the about page and 2 + 2 is %d", sum))
}

// addValues adds two integers and return the sum
func AddValues(x, y int) int {
	return x + y
}

func Divide(w http.ResponseWriter, r *http.Request) {
	f, err := divideValues(100.0, 10.0)
	if err != nil {
		fmt.Fprintf(w, "Cannot divide by zero.")
		return
	}
	fmt.Fprintf(w, fmt.Sprintf("%f divided by %f is %f", 100.0, 10.0, f))
}

func divideValues(x, y float32) (float32, error) {
	if y < 0 {
		err := errors.New("Cannot divide by zero.")
		return 0, err
	}
	return x / y, nil
}
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	fmt.Println(fmt.Sprintf("Starting server on port: %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
