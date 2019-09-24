package main

import "fmt"

var x = "har har"

func main() {
	fmt.Println(x)
	//x = 6 // Can't do that, x is a string type variable
	x = "yuk yuk"
	fmt.Println(x) // That works fine, new value is a string too.
}
