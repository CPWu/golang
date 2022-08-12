package main

import "log"

func main() {
  myMap := make(map[string]string)
  
  myMap["dog"] = "Samson"
  myMap["other-dog"] = "Cassie"
  myMap["dog"] = "Fido"


  log.Println(myMap["dog"])
  log.Println(myMap["other-dog"])
  log.Println(myMap["dog"])
  
  // Slices
  var mySlice []string

  mySlice = append(mySlice, "Chun")
  mySlice = append(mySlice, "Miti")

  numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

  log.Println(numbers)
  log.Println(mySlice)
  log.Println(numbers[0:3])
}
