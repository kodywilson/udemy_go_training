package main

import "fmt"

type foo int
type bar string

func main() {
	var myAge foo
	myAge = 44
	fmt.Printf("%T %v \n", myAge, myAge)
	var myName bar
	myName = "Kody"
	fmt.Printf("%T %v \n", myName, myName)
}
