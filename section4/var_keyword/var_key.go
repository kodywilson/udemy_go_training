package main

import "fmt"

var x = 3

func main() {
	fmt.Println(x) // This works, x is in scope.
	y := 4
	fmt.Println(y) // Ok, y is in scope here.
	foo()
}

func foo() {
	fmt.Println(x) // This will work too.
}
