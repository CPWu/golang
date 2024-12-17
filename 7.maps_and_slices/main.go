package main

import (
	"log"
	"sort"
)

// Concepts
// Maps and Slices
// Maps by default are not sorted, while initial implementation this was the case.

type User struct {
	FirstName string
	LastName  string
}

func main() {
	// ** Maps **
	// Declare First then Initialize
	var myOtherMap map[string]string
	myOtherMap = map[string]string{"cat": "Luna"}

	// Declare and Initialize in One Line
	myOtherOtherMap := map[string]string{"fish": "Nemo"}

	// Delcare and Make Empty Map, then Add Values
	var myOtherOtherOtherMap map[string]string
	myOtherOtherOtherMap = make(map[string]string)
	myOtherOtherOtherMap["bird"] = "Tweety"

	// Most Common - shorthand
	myMap := make(map[string]string)
	myMap["dog"] = "Samson"
	myMap["other-dog"] = "Cassie"
	// Overwriting value

	myMap["dog"] = "Fido"

	// Int Map
	myIntMap := make(map[string]int)
	myIntMap["First"] = 1
	myIntMap["Second"] = 2

	// User defined data type map
	myUserMap := make(map[string]User)
	me := User{
		FirstName: "Trevor",
		LastName:  "Sawler",
	}
	myUserMap["me"] = me

	// Output
	log.Println("Dog's Name:", myMap["dog"])
	log.Println("Other Dog's Name:", myMap["other-dog"])
	log.Println("Cat's Name:", myOtherMap["cat"])
	log.Println("Fish's Name:", myOtherOtherMap["fish"])
	log.Println("Bird's Name:", myOtherOtherOtherMap["bird"])

	log.Println("My int map:", myIntMap["First"], myIntMap["Second"])

	log.Println("My user map:", myUserMap["me"].FirstName)

	// ** Slices **
	// String Slice
	var mySlice []string
	mySlice = append(mySlice, "Chun")
	mySlice = append(mySlice, "Miti")
	mySlice = append(mySlice, "Teemo")

	// Int Slice
	numbers := []int{100, 2, 30, 4, 50, 6, 7, 8, 9, 10}
	sort.Ints(numbers)

	log.Println(numbers)
	log.Println("My slice:", mySlice)
	log.Println("Int slice (sorted) in range:", numbers[0:3])
}
