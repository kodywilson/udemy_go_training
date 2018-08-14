package main

import "fmt"

func main() {

  var myGreeting = make(map[string]string)
  myGreeting["Terry"] = "Hola!"
  myGreeting["Ernesto"] = "Good morning!"

  fmt.Println(myGreeting)
}
