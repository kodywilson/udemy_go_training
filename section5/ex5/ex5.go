package main

import "fmt"

type numby int

var x numby
var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
	y = int(x)
	fmt.Printf("Type: %T and Value: %v\n", y, y)
}
