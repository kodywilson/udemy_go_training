package main

import "fmt"

func main() {
	greet("Kody", "Wilson")
}

func greet(fname, lname string) {
	fmt.Println(fname, lname)
}

// same two parameters, both are strings so using shorthand
