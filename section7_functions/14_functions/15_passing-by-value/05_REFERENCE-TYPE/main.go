package main

import "fmt"

func main() {
  m := make(map[string]int)
  changeMe(m)
  fmt.Println(m["Kody"]) // 41
}

func changeMe(z map[string]int) {
  z["Kody"] = 41
}
