package main

import "fmt"

func main() {
  x := 33
  if x < 60 || false {
    fmt.Println("This ran")
  }
}
