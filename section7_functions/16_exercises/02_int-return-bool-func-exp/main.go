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
	fmt.Println("You passed 100. Now I divide by 2 and tell you if even.")
	fmt.Println(dividebool(100))
}
