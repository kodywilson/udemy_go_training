package main

import "fmt"

func main() {
	x := (true && false) || (false && true) || !(false && false)
	fmt.Println(x)
}
