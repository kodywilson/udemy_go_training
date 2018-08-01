package main

import "fmt"

func main() {
	x := (true && false) || (false && true) || !(false && false)
	y := true && !false
	z := true && false
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
