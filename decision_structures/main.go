package main

import "log"

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
    log.Println("cat is: ", cat)
  } else {
    log.Println("cat is: ", cat)
  }

  // Decision Structure - Switch

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
