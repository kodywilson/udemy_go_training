package main

import "fmt"

func main() {

  a := []string{"Monday", "Tuesday"}
  b := []string{"Wednesday", "Thursday", "Friday"}

  c := append(a, b...)
  fmt.Println(c)

  // Deleting is weird
  c = append(c[:2], c[3:]...)
  fmt.Println(c)
}
