package main

import "log"

// Concepts
// Decision Structures

func main() {
	// Decision Structure - If/Else
	var isTrue bool
	isTrue = true

	if isTrue {
		log.Println("isTrue is", isTrue)
	} else {
		log.Println("isTrue is", isTrue)
	}

	cat := "cat"

	if cat == "cat" {
		log.Println("cat is:", cat)
	} else {
		log.Println("cat is:", cat)
	}

	// Decision Structure - Switch
	// Golang specific: Switch statement does not continue after a match.
	switchCase := "dog"

	switch switchCase {
	case "cat":
		log.Println("You have a cat.")
	case "dog":
		log.Println("You have a dog.")
	case "turtle":
		log.Println("You have a turtle.")
	default:
		log.Println("You dont have any animals.")
	}
}
