package main

import "fmt"

func main() {
	dividebool := func(x int) (int, bool) {
		var even bool
		if x%2 == 0 {
			even = true
		}
		return x / 2, even
	}
	var x int
	fmt.Println("Please enter a whole number (no decimals): ")
	_, err := fmt.Scanf("%d", &x)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("You passed", x, "- Now I divide by 2 and tell you if even.")
	fmt.Println(dividebool(x))
}
