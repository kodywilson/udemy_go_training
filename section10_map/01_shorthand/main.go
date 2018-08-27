package main

import "fmt"

func main() {

  myGreeting := make(map[string]string)
  myGreeting["Terry"] = "Hola!"
  myGreeting["Ernesto"] = "Good morning!"

  fmt.Println(myGreeting)

  myInts := make(map[int]int)
  myInts[0] = 0
  myInts[1] = 1
  fmt.Println(myInts)
} 
