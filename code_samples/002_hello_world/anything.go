package main

import "fmt"

func main() {
	fmt.Println("Hello everyone, this is a lot of fun.")
	foo()
	fmt.Println("Now, I am out of foo!")

	x := 3

	for i := 0; i < 20; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	fmt.Println(x)
}

func foo() {
	fmt.Println("I'm in foo.")
	bar()
}

func bar() {
	fmt.Println("Welcome to bar")
	city()
}

func city() {
	fmt.Println("it's lovely here in the city.")
}

// a comment
