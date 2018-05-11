package main

import "fmt"

func main() {
	y := "Tinker"
	fmt.Println(x)
	fmt.Println(y)
	foo()
}

func foo() {
	y := "Waddlegump"
	fmt.Println(x)
	fmt.Println(y)
}

// order doesn't matter for package level scope
var x string = "Dorwitz"
