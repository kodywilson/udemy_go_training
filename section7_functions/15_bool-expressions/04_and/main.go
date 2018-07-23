package main

import "fmt"

func main() {
  x := 33
  if x < 60 && false {
    fmt.Println("This did not run")
  }
  if x < 60 && true {
    fmt.Println("This ran")
  }
}
