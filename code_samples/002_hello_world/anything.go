package main

import "fmt"

func main() {
	fmt.Println("Hello everyone, this is a lot of fun.")
	foo()
	fmt.Println("something more")

	x := 3

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	fmt.Println(x)
}

func foo() {
	fmt.Println("I'm in foo.")
}

// a comment
