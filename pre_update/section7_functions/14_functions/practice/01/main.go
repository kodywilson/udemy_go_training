package main

import "fmt"

func calc(x, y int) int {
	return x + y
}

func main() {
	fmt.Println("This is only a test.")
	fmt.Println("Jk, let's add some numbers.")
	var x int
	fmt.Println("Please enter a whole number (no decimals): ")
	_, err := fmt.Scanf("%d", &x)
	if err != nil {
		fmt.Println(err)
	}
	var y int
	fmt.Println("Please enter a whole number (no decimals): ")
	_, err = fmt.Scanf("%d", &y)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(x, "+", y, "=", calc(x, y))
	fmt.Println("Hello!")
	for i := 0; i < 3; i++ {
		fmt.Println("Bye!")
	}
}
