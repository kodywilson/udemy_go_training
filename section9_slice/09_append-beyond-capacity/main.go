package main

import "fmt"

func main() {

  greeting := make([]string, 3, 5)
  // 3 is length
  // 5 is capacity

  greeting[0] = "Good morning!"
  greeting[1] = "Bonjour!"
  greeting[2] = "Buenos dias!"
  greeting = append(greeting, "Suprabadham!")
  greeting = append(greeting, "Zao an")
  greeting = append(greeting, "Ohayou gozaimasu")
  greeting = append(greeting, "Gidday!")

  fmt.Println(greeting[6])
  fmt.Println(len(greeting))
  fmt.Println(cap(greeting))

}
