package main

import "fmt"

var x int = 42

func main() {
	y := 43
	fmt.Println(x)
	fmt.Println(y)
	foo()
}

func foo() {
	y := 41
	fmt.Println(x)
	fmt.Println(y)
}
