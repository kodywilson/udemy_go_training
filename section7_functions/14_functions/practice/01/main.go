package main

import "fmt"

func calc(x, y int) int {
  return x + y
}

func main() {
  fmt.Println("This is only a test.")
  x := 100
  y := 63
  fmt.Println("Jk, let's add some numbers.")
  fmt.Println(x, "+", y, "=", calc(x, y))
  fmt.Println("Hello!")
  for i := 0; i < 10; i++ {
    fmt.Println("Bye!")
  }
}
