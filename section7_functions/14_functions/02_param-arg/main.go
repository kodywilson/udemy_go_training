package main

import "fmt"

func main() {
	greet("Kody")
	greet("Heidi")
}

func greet(name string) {
	fmt.Println(name)
}

// greet is declared with a parameter
// when calling greet, pass an argument
