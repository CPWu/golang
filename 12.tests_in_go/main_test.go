package main

import "testing"

// Run `go test -cover for Test Coverage`
// OR
// `go test -coverprofile=coverage.out && go tool cover -html=coverage.out`
// Table Tests
// tests is equivalent to slice of structs, each struct defines a test case to run.
var tests = []struct {
	name     string
	dividend float32
	divisor  float32
	expected float32
	isErr    bool
}{
	{"valid-data", 100.0, 10.0, 10.0, false},
	{"invalid-data", 100.0, 0.0, 0.0, true},
	{"expect-five", 50.0, 10.0, 5.0, false},
}

// To write a test in go, you need to give it a name and it has to start with the word test.
// Run tests by running `go test` or `go test -v`
func TestDivision(t *testing.T) {
	// Range over test variable
	for _, tt := range tests {
		got, err := divide(tt.dividend, tt.divisor)
		if tt.isErr {
			if err == nil {
				t.Error("Expected an error but did not get one.")
			}
		} else {
			if err != nil {
				t.Error("Did not expect an error but got one.", err.Error())
			}
		}

		if got != tt.expected {
			t.Errorf("Expected %f but got %f", tt.expected, got)
		}
	}
}

// func TestDivide(t *testing.T) {
// 	_, err := divide(10.0, 1.0)

// 	if err != nil {
// 		t.Error("Got an error when we should not have.")
// 	}
// }

// func TestBadDivide(t *testing.T) {
// 	_, err := divide(10.0, 0)

// 	if err == nil {
// 		t.Error("Did not get an error.")
// 	}
// }
