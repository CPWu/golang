package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "Hello World")
		if err != nil {
			// The n and err return values from Fprintf are
			// those returned by the underlying io.Writer.
			fmt.Fprintf(os.Stderr, "Fprintf: %v\n", err)
		}
		fmt.Printf("%d bytes written.\n", n)
	})

	_ = http.ListenAndServe(":8080", nil)

}
