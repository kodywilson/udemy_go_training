package main

import "fmt"

func main() {
	example := []string{}
	examples := [][]string{}
	fmt.Println("An example with shorthand:")
	fmt.Println(example)
	fmt.Println(examples)
	fmt.Println(example == nil)
	// It isn't nil, but you have to use append to start adding to it

	var example2 []string
	var examples2 [][]string
	fmt.Println("An example with var:")
	fmt.Println(example2)
	fmt.Println(examples2)
	fmt.Println(example2 == nil) // it is nil

	example3 := make([]string, 35)
	examples3 := make([]string, 35)
	fmt.Println("An example with make:")
	fmt.Println(example3)
	fmt.Println(examples3)
	fmt.Println(example3 == nil) // not nil and can be referenced by index
}
